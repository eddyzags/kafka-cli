package client

import (
	"strings"
)

type Client struct {
	URL string
}

func NewClient(apiURL string) *Client {
	if !strings.HasPrefix(apiURL, "http://") {
		apiURL = "http://" + apiURL
	}

	return &Client{
		URL: apiURL,
	}
}
