package rest

import (
	"fmt"
	"net/http"
)

type BaseClient struct {
	URL        string
	HTTPClient *http.Client
}

func (c *BaseClient) SendRequest(req *http.Request) (*http.Response, error) {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	response, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request. error: %w", err)
	}

	return response, nil
}
