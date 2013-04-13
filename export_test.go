package gsb

import (
	"io"
	"net/url"
)

// Exporting functions for testing

func ListURL(u *url.URL, key string) *url.URL {
	return urlForListRequest(u, key)
}

func (c *Client) ParseListResponse(l []string, r io.Reader) ([]string, error) {
	return c.parseListResponse(l, r)
}
