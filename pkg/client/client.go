package client

type ApodClient interface{}

type Client struct {
	ApodClient
	url string
}

func NewClient(url string) *Client {
	return &Client{url: url}
}
