package client

import (
	"bufio"
	"context"
	"errors"
	"net/http"
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

func (c *Client) NewRequest(ctx context.Context, method, url string) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, method, url, nil)
}

func isFailureStatusCode(resp *http.Response) bool {
	return resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest
}
