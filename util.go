package hqgohttp

import (
	"io"
)

type ContextOverride string

const (
	RetryMax ContextOverride = "retry-max"
)

// getLength returns length of a Reader efficiently
func getLength(x io.ReadCloser) (int64, error) {
	length, err := io.Copy(io.Discard, x)

	return length, err
}

func getReusableBodyandContentLength(rawBody interface{}) (io.ReadCloser, int64, error) {
	var bodyReader io.ReadCloser

	var contentLength int64

	if rawBody != nil {
		switch body := rawBody.(type) {
		// If they gave us a function already, great! Use it.
		case io.ReadCloser:
			bodyReader = body
		case *io.ReadCloser:
			bodyReader = *body
		}
	}

	if bodyReader != nil {
		var err error

		contentLength, err = getLength(bodyReader)
		if err != nil {
			return nil, 0, err
		}
	}

	return bodyReader, contentLength, nil
}
