package client

import (
	"net/http"
	"net/url"
	"strings"
)

// Get issues a GET to the specified URL.
func Get(URL string) (*http.Response, error) {
	return DefaultHTTPClient.Get(URL)
}

// Get is a convenience helper for doing simple GET requests.
func (c *Client) Get(URL string) (*http.Response, error) {
	req, err := NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

// Head issues a HEAD to the specified URL.
func Head(URL string) (*http.Response, error) {
	return DefaultHTTPClient.Head(URL)
}

// Head is a convenience method for doing simple HEAD requests.
func (c *Client) Head(URL string) (*http.Response, error) {
	req, err := NewRequest(http.MethodHead, URL, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

// Post issues a POST to the specified URL.
func Post(URL, bodyType string, body interface{}) (*http.Response, error) {
	return DefaultHTTPClient.Post(URL, bodyType, body)
}

// Post is a convenience method for doing simple POST requests.
func (c *Client) Post(URL, bodyType string, body interface{}) (*http.Response, error) {
	req, err := NewRequest(http.MethodPost, URL, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", bodyType)

	return c.Do(req)
}

// PostForm issues a POST to the specified URL, with data's keys and values
func PostForm(URL string, data url.Values) (*http.Response, error) {
	return DefaultHTTPClient.PostForm(URL, data)
}

// PostForm is a convenience method for doing simple POST operations using
// pre-filled url.Values form data.
func (c *Client) PostForm(URL string, data url.Values) (*http.Response, error) {
	return c.Post(URL, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}
