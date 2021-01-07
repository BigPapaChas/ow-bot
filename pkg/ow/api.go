package ow

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// Client will hold the information for consuming from ow-api
type Client struct {
	BaseURL *url.URL

	httpClient *http.Client
}

// NewClient returns a pointer to a new ow-api Client where the httpClient can be configured
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	c := &Client{
		httpClient: httpClient,
	}
	return c
}

// NewDefaultClient returns a pointer to new ow-api Client that uses the default httpClient
func NewDefaultClient() *Client {
	return NewClient(http.DefaultClient)
}

// GetProfile performs a GET on ow-api's /profile endpoint and returns a Profile struct
func (c *Client) GetProfile() (*Profile, error) {
	req, err := c.newRequest("GET", "./profile", nil)
	if err != nil {
		return nil, err
	}

	p := &Profile{}
	_, err = c.do(req, p)
	return p, err
}

// newRequest will create and return a pointer to a http.Request object
func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	return req, nil
}

// do performs a http request and decodes the response body into the passed interface
func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
