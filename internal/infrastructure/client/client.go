package client

import (
	"bufio"
	"context"
	"errors"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	client  *http.Client
	request *http.Request

	err error

	config ClientConfig
}

var (
	defaultTimeout = 30 * time.Second
)

func NewClient() *Client {
	config := DefaultConfig()
	return NewClientWithConfig(config)
}

func NewClientWithConfig(config ClientConfig) *Client {
	return &Client{
		config: config,
	}
}

func (c *Client) SendRequestStream(req *http.Request) (*StreamReader, error) {
	req.Header.Set("Content-type", "text/event-stream")
	req.Header.Set("Accept-Charset", "utf-8")
	req.Header.Set("Connection", "keep-alive")

	response, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return new(StreamReader), err
	}

	if isFailureStatusCode(response) {
		return new(StreamReader), errors.New("request failed")
	}

	return &StreamReader{
		response: response,
		reader:   bufio.NewReader(response.Body),
	}, nil
}

func (c *Client) Do(request *http.Request) (*http.Response, error) {
	return c.config.HTTPClient.Do(request)
}

func (c *Client) PostForm(request *http.Request, values url.Values) (*http.Response, error) {
	return c.config.HTTPClient.PostForm(request.RequestURI, values)
}

func (c *Client) HttpClient() *http.Client {
	return c.config.HTTPClient
}

func (c *Client) NewRequest(ctx context.Context, method, url string) (*http.Request, error) {
	request, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return request, err
	}
	c.request = request
	return request, err
}

func isFailureStatusCode(resp *http.Response) bool {
	return resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest
}

func (c *Client) SetHeader(key, value string) {
	c.request.Header.Set(key, value)
}
