package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type JsonHttpClient struct {
	BaseUrl string
	client  *http.Client
}

func (c *JsonHttpClient) fullUrl(path string, params map[string]string) (string, error) {
	stringUrl, err := url.JoinPath(c.BaseUrl, path)
	if err != nil {
		return "", fmt.Errorf("failed to join url: %w", err)
	}
	u, err := url.Parse(stringUrl)
	if err != nil {
		return "", fmt.Errorf("failed to parse requested url: %w", err)
	}
	q := u.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}

func (c *JsonHttpClient) Request(method string, path string, headers map[string][]string, params map[string]string, body []byte) (respBody []byte, err error) {
	url, err := c.fullUrl(path, params)
	if err != nil {
		return nil, err
	}
	bodyReader := bytes.NewReader(body)
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to assemble http request: %w", err)
	}
	req.Header = headers
	headers["Content-Type"] = []string{"application/json"}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to request: %w", err)
	}
	respBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	return respBody, nil
}
