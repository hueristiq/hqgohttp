// This file contains code focusing on the construction and configuration of an HTTP client that provides functionality like automatic retries, backoff strategies, and logging hooks.
package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	dac "github.com/Mzack9999/go-http-digest-auth-client"
	"golang.org/x/net/http2"
)

// Client represents the main HTTP client. It is used to make HTTP requests and
// adds additional functionality like automatic retries to tolerate minor outages.
type Client struct {
	// HTTPClient is the internal HTTP client (http1x + http2 via connection upgrade upgrade).
	HTTPClient *http.Client
	// HTTPClient is the internal HTTP client configured to fallback to native http2 at transport level
	HTTPClient2 *http.Client
	// RequestLogHook allows a user-supplied function to be called
	// before each retry.
	RequestLogHook RequestLogHook
	// ResponseLogHook allows a user-supplied function to be called
	// with the response from each HTTP request executed.
	ResponseLogHook ResponseLogHook
	// ErrorHandler specifies the custom error handler to use, if any
	ErrorHandler ErrorHandler
	// CheckRetry specifies the policy for handling retries, and is called
	// after each request
	CheckRetry CheckRetry
	// Backoff specifies the policy for how long to wait between retries
	Backoff Backoff

	requestCounter uint32

	options Options
}

// setKillIdleConnections sets the kill idle conns switch in two scenarios
//  1. If the http.Client has settings that require us to do so.
//  2. The user has enabled it by default, in which case we have nothing to do.
func (c *Client) setKillIdleConnections() {
	if c.HTTPClient != nil || !c.options.KillIdleConn {
		if b, ok := c.HTTPClient.Transport.(*http.Transport); ok {
			c.options.KillIdleConn = b.DisableKeepAlives || b.MaxConnsPerHost < 0
		}
	}
}

// Do wraps calling an HTTP method with retries.
func (c *Client) Do(req *Request) (*http.Response, error) {
	var resp *http.Response

	var err error

	// Create a main context that will be used as the main timeout
	mainCtx, cancel := context.WithTimeout(context.Background(), c.options.Timeout)

	defer cancel()

	retryMax := c.options.RetryMax

	if ctxRetryMax := req.Context().Value(RetryMax); ctxRetryMax != nil {
		if maxRetriesParsed, ok := ctxRetryMax.(int); ok {
			retryMax = maxRetriesParsed
		}
	}

	for i := 0; ; i++ {
		// request body can be read multiple times
		// hence no need to rewind it
		if c.RequestLogHook != nil {
			c.RequestLogHook(req.Request, i)
		}

		if req.hasAuth() && req.Auth.Type == DigestAuth {
			digestTransport := dac.NewTransport(req.Auth.Username, req.Auth.Password)
			digestTransport.HTTPClient = c.HTTPClient
			resp, err = digestTransport.RoundTrip(req.Request)
		} else {
			// Attempt the request with standard behavior
			resp, err = c.HTTPClient.Do(req.Request)
		}

		// Check if we should continue with retries.
		checkOK, checkErr := c.CheckRetry(req.Context(), resp, err)

		// if err is equal to missing minor protocol version retry with http/2
		if err != nil && strings.Contains(err.Error(), "net/http: HTTP/1.x transport connection broken: malformed HTTP version \"HTTP/2\"") {
			resp, err = c.HTTPClient2.Do(req.Request)
			checkOK, checkErr = c.CheckRetry(req.Context(), resp, err)
		}

		if err != nil {
			// Increment the failure counter as the request failed
			req.Metrics.Failures++
		} else if c.ResponseLogHook != nil {
			// Call this here to maintain the behavior of logging all requests,
			// even if CheckRetry signals to stop.

			// Call the response logger function if provided.
			c.ResponseLogHook(resp)
		}

		// Now decide if we should continue.
		if !checkOK {
			if checkErr != nil {
				err = checkErr
			}

			c.closeIdleConnections()

			return resp, err
		}

		// We do this before drainBody beause there's no need for the I/O if
		// we're breaking out
		remain := retryMax - i
		if remain <= 0 {
			break
		}

		// Increment the retries counter as we are going to do one more retry
		req.Metrics.Retries++

		// We're going to retry, consume any response to reuse the connection.
		if err == nil && resp != nil {
			c.drainBody(req, resp)
		}

		// Wait for the time specified by backoff then retry.
		// If the context is cancelled however, return.
		wait := c.Backoff(c.options.RetryWaitMin, c.options.RetryWaitMax, i, resp)

		// Exit if the main context or the request context is done
		// Otherwise, wait for the duration and try again.
		// use label to explicitly specify what to break
		select {
		case <-mainCtx.Done(): // Do nothing; it will break out of the select block by default.
		case <-req.Context().Done():
			c.closeIdleConnections()

			return nil, req.Context().Err()
		case <-time.After(wait): // Do nothing; it will continue after the wait duration.
		}
	}

	if c.ErrorHandler != nil {
		c.closeIdleConnections()

		return c.ErrorHandler(resp, err, retryMax+1)
	}

	// By default, we close the response body and return an error without
	// returning the response
	if resp != nil {
		resp.Body.Close()
	}

	c.closeIdleConnections()

	return nil, fmt.Errorf("%s %s giving up after %d attempts: %w", req.Method, req.URL, retryMax+1, err)
}

// Try to read the response body so we can reuse this connection.
func (c *Client) drainBody(req *Request, resp *http.Response) {
	_, err := io.Copy(io.Discard, io.LimitReader(resp.Body, c.options.RespReadLimit))
	if err != nil {
		req.Metrics.DrainErrors++
	}

	resp.Body.Close()
}

func (c *Client) closeIdleConnections() {
	if c.options.KillIdleConn {
		requestCounter := atomic.LoadUint32(&c.requestCounter)
		if requestCounter < closeConnectionsCounter {
			atomic.AddUint32(&c.requestCounter, 1)
		} else {
			atomic.StoreUint32(&c.requestCounter, 0)
			c.HTTPClient.CloseIdleConnections()
		}
	}
}

// Options represents configuration options for the client
type Options struct {
	// RetryWaitMin is the minimum time to wait for retry
	RetryWaitMin time.Duration
	// RetryWaitMax is the maximum time to wait for retry
	RetryWaitMax time.Duration
	// Timeout is the maximum time to wait for the request
	Timeout time.Duration
	// RetryMax is the maximum number of retries
	RetryMax int
	// RespReadLimit is the maximum HTTP response size to read for
	// connection being reused.
	RespReadLimit int64
	// Verbose specifies if debug messages should be printed
	Verbose bool
	// KillIdleConn specifies if all keep-alive connections gets killed
	KillIdleConn bool
	// Custom CheckRetry policy
	CheckRetry CheckRetry
	// Custom Backoff policy
	Backoff Backoff
	// NoAdjustTimeout disables automatic adjustment of HTTP request timeout
	NoAdjustTimeout bool
	// Custom http client
	HTTPClient *http.Client
}

const closeConnectionsCounter = 100

// DefaultOptionsSpraying contains the default options for host spraying
// scenarios where lots of requests need to be sent to different hosts.
var DefaultOptionsSpraying = &Options{
	RetryWaitMin:    1 * time.Second,
	RetryWaitMax:    30 * time.Second,
	Timeout:         30 * time.Second,
	RetryMax:        5,
	RespReadLimit:   4096,
	KillIdleConn:    true,
	NoAdjustTimeout: true,
}

// DefaultOptionsSingle contains the default options for host bruteforce
// scenarios where lots of requests need to be sent to a single host.
var DefaultOptionsSingle = &Options{
	RetryWaitMin:    1 * time.Second,
	RetryWaitMax:    30 * time.Second,
	Timeout:         30 * time.Second,
	RetryMax:        5,
	RespReadLimit:   4096,
	KillIdleConn:    false,
	NoAdjustTimeout: true,
}

// DefaultHTTPClient is the http client with DefaultOptionsSingle options.
var DefaultHTTPClient *Client

func init() {
	DefaultHTTPClient = New(DefaultOptionsSingle)
}

// New creates a new client instance based on provided options.
// It configures the internal HTTP clients, sets up HTTP/2 for the second client,
// applies retry and backoff policies, and Adjusts client timeouts and
// other settings based on the provided options.
func New(options *Options) (client *Client) {
	client = &Client{
		options: *options,
	}

	client.HTTPClient = DefaultClient()

	if options.HTTPClient != nil {
		client.HTTPClient = options.HTTPClient
	}

	client.HTTPClient2 = DefaultClient()

	if err := http2.ConfigureTransport(client.HTTPClient2.Transport.(*http.Transport)); err != nil {
		return
	}

	client.CheckRetry = DefaultRetryPolicy() //nolint:bodyclose // To be refactored

	if options.CheckRetry != nil {
		client.CheckRetry = options.CheckRetry
	}

	client.Backoff = DefaultBackoff() //nolint:bodyclose // To be refactored

	if options.Backoff != nil {
		client.Backoff = options.Backoff
	}

	// add timeout to clients
	if options.Timeout > 0 {
		client.HTTPClient.Timeout = options.Timeout
		client.HTTPClient2.Timeout = options.Timeout
	}

	// if necessary adjusts per-request timeout proportionally to general timeout (30%)
	if options.Timeout > time.Second*15 && options.RetryMax > 1 && !options.NoAdjustTimeout {
		client.HTTPClient.Timeout = time.Duration(options.Timeout.Seconds()*0.3) * time.Second
	}

	client.setKillIdleConnections()

	return
}
