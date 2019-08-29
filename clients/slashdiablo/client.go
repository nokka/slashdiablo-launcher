package slashdiablo

import (
	"fmt"
	"io"
	"net/http"
)

// Client ...
type Client struct {
	address string
}

// GetFile will the file by the given path in the repository set on the service.
func (c *Client) GetFile(filePath string) (io.ReadCloser, error) {
	addr := fmt.Sprintf("%s/%s", c.address, filePath)
	fmt.Println("CLIENT ADDRESS")
	fmt.Println(addr)
	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

// NewClient returns a new client with all dependencies setup.
func NewClient(address string) Client {
	return Client{
		address: address,
	}
}
