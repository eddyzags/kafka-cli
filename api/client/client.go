package client

type Client struct {
	URL string
}

func NewClient(apiURL string) *Client {
	return &Client{
		URL: apiURL,
	}
}
