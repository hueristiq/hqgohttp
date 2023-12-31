package hqgohttp

// This file contains set of Go functions that focuses on handling HTTP request retries based on specific conditions.

import (
	"context"
	"crypto/x509"
	"errors"
	"net/http"
	"net/url"
	"regexp"
)

var (
	// A regular expression to match the error returned by net/http when the
	// configured number of redirects is exhausted. This error isn't typed
	// specifically so we resort to matching on the error string.
	redirectsErrorRegex = regexp.MustCompile(`stopped after \d+ redirects\z`)

	// A regular expression to match the error returned by net/http when the
	// scheme specified in the URL is invalid. This error isn't typed
	// specifically so we resort to matching on the error string.
	schemeErrorRegex = regexp.MustCompile(`unsupported protocol scheme`)
)

// CheckRetry specifies a policy for handling retries. It is called
// following each request with the response and error values returned by
// the http.Client. If CheckRetry returns false, the Client stops retrying
// and returns the response to the caller. If CheckRetry returns an error,
// that error value is returned in lieu of the error from the request. The
// Client will close any response body when retrying, but if the retry is
// aborted it is up to the CheckRetry callback to properly close any
// response body before returning.
type CheckRetry func(ctx context.Context, resp *http.Response, err error) (bool, error)

// DefaultRetryPolicy provides a default callback for client.CheckRetry, which
// will retry on connection errors and server errors.
func DefaultRetryPolicy() func(ctx context.Context, resp *http.Response, err error) (bool, error) {
	return CheckRecoverableErrors
}

// HostSprayRetryPolicy provides a callback for client.CheckRetry, which
// will retry on connection errors and server errors.
func HostSprayRetryPolicy() func(ctx context.Context, resp *http.Response, err error) (bool, error) {
	return CheckRecoverableErrors
}

// CheckRecoverableErrors checks if an error is recoverable and decides
// whether to retry the request. The conditions it checks are:
// 1. If the context has been canceled or its deadline has been exceeded, it doesn't retry.
// 2. If the error is related to too many redirects or an unsupported protocol scheme, it doesn't retry.
// 3. If the error is due to a TLS certificate verification failure (specifically an unknown authority error), it doesn't retry.
// If none of the above conditions are met, it considers the error as likely recoverable and decides to retry.
func CheckRecoverableErrors(ctx context.Context, _ *http.Response, err error) (bool, error) {
	// do not retry on context.Canceled or context.DeadlineExceeded
	if ctx.Err() != nil {
		return false, ctx.Err()
	}

	if err == nil {
		return false, nil
	}

	var urlErr *url.Error

	if errors.As(err, &urlErr) {
		// Don't retry if the error was due to too many redirects.
		// Don't retry if the error was due to an invalid protocol scheme.
		// Don't retry if the error was due to TLS cert verification failure.
		if isRedirectError(urlErr) || isSchemeError(urlErr) || isUnknownAuthorityError(urlErr) {
			return false, nil
		}
	}

	// The error is likely recoverable so retry.
	return true, nil
}

func isRedirectError(err *url.Error) bool {
	return redirectsErrorRegex.MatchString(err.Error())
}

func isSchemeError(err *url.Error) bool {
	return schemeErrorRegex.MatchString(err.Error())
}

func isUnknownAuthorityError(err *url.Error) bool {
	var authorityErr x509.UnknownAuthorityError

	return errors.As(err.Err, &authorityErr)
}
