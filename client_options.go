package hqgohttp

// This Go code defines a structure Options to represent configuration options for an HTTP client,
// and two variables, DefaultOptionsSpraying and DefaultOptionsSingle, which are instances of
// Options with default values for two different scenarios.

import (
	"net/http"
	"time"
)

// Options represents configuration fields to customize the behavior of the HTTP client
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

// DefaultOptionsSingle is an instance of Options with default values suitable for
// "host brute force" scenarios, where lots of requests need to be sent to a single
// host. For example, it sets KillIdleConn to false to allow keep-alive connections,
// as they can improve performance when connecting repeatedly to the same host.
var DefaultOptionsSingle = &Options{
	RetryWaitMin:    1 * time.Second,
	RetryWaitMax:    30 * time.Second,
	Timeout:         30 * time.Second,
	RetryMax:        5,
	RespReadLimit:   4096,
	KillIdleConn:    false,
	NoAdjustTimeout: true,
}

// DefaultOptionsSpraying is an instance of Options with default values suitable for
// "host spraying" scenarios, where lots of requests need to be sent to different hosts.
// For example, it sets KillIdleConn to true to kill all keep-alive connections,
// as they are not useful when connecting to many different hosts.
var DefaultOptionsSpraying = &Options{
	RetryWaitMin:    1 * time.Second,
	RetryWaitMax:    30 * time.Second,
	Timeout:         30 * time.Second,
	RetryMax:        5,
	RespReadLimit:   4096,
	KillIdleConn:    true,
	NoAdjustTimeout: true,
}
