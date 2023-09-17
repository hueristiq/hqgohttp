package hqgohttp

// This file contains utility functions to create HTTP clients and transports.

import (
	"net"
	"net/http"
	"runtime"
	"time"
)

// DefaultHTTPTransport returns a new http.Transport with similar default values to
// http.DefaultTransport, but with idle connections and keepalives disabled.
// It does this by first creating a transport with pooled connections
// (by calling DefaultHTTPPooledTransport) and then setting DisableKeepAlives
// to true and MaxIdleConnsPerHost to -1.
func DefaultHTTPTransport() (transport *http.Transport) {
	transport = DefaultHTTPPooledTransport()
	transport.DisableKeepAlives = true
	transport.MaxIdleConnsPerHost = -1

	return
}

// DefaultHTTPPooledTransport returns a new http.Transport with similar default
// values to http.DefaultTransport, but with a custom configuration that is
// suitable for transports that will be reused for the same hosts. It sets various
// fields of the http.Transport struct, such as Proxy, DialContext, MaxIdleConns,
// IdleConnTimeout, TLSHandshakeTimeout, ExpectContinueTimeout, ForceAttemptHTTP2, and
// MaxIdleConnsPerHost.
//
// Do not use this for transient transports as it can leak file descriptors over
// time. Only use this for transports that will be re-used for the same host(s).
func DefaultHTTPPooledTransport() (transport *http.Transport) {
	transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		ForceAttemptHTTP2:     true,
		MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
	}

	return
}

// DefaultHTTPClient returns a new http.Client with similar default values to
// http.Client, but with a non-shared transport, idle connections disabled, and
// keep-alives disabled. It does this by setting the Transport field of the http.Client
// struct to the transport returned by DefaultHTTPTransport.
func DefaultHTTPClient() *http.Client {
	return &http.Client{
		Transport: DefaultHTTPTransport(),
	}
}

// DefaultPooledClient returns a new http.Client with similar default values to
// http.Client, but with a shared transport. It sets the Transport field of the
// http.Client struct to the transport returned by DefaultHTTPPooledTransport.
//
// Do not use this function for transient clients as it can leak file descriptors
// over time. Only use this for clients that will be re-used for the same host(s).
func DefaultPooledClient() *http.Client {
	return &http.Client{
		Transport: DefaultHTTPPooledTransport(),
	}
}
