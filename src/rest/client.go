package rest

import (
	"log"
	"net/url"
	"net/http"
)

type Client struct {
	client *http.Client
}

func NewClient(proxy string) *Client {
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		log.Fatalf("Failed to parse proxy URL: %v", err)
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	return &Client{
		client: &http.Client{
			Transport: transport,
		},
	}
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		log.Printf("Failed to do request: %v", err)
		return nil, err
	}
	return resp, nil
}