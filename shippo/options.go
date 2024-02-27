package shippo

import (
	"fmt"
	"net/http"
	"net/url"
)

// ClientOption describes a functional parameter for the client constructor.
type ClientOption func(*client) error

// WithHttpClient set the http client.
func WithHttpClient(cl *http.Client) ClientOption {
	return func(c *client) error {
		if cl == nil {
			return fmt.Errorf("http client cannot be nil")
		}

		c.client = cl
		return nil
	}
}

func WithBaseUrl(address string) ClientOption {
	return func(c *client) error {
		baseUrl, err := url.Parse(address)
		if err != nil {
			return err
		}

		c.baseURL = baseUrl
		return nil
	}
}
