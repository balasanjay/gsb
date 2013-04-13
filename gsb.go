package gsb

import (
	"bufio"
	"io"
	"net/http"
	"net/url"
	"os"
)

const (
	API_KEY_ENV_VAR  = "SAFE_BROWSE_API_KEY"
	CLIENT_VERSION   = "0.0.1"
	PROTOCOL_VERSION = "2.2"
	API_HOST         = "safebrowsing.clients.google.com"
	API_SCHEME       = "https"
	API_LIST_PATH    = "/safebrowsing/list"
)

type Client struct {
	Key        string
	HttpClient *http.Client
}

var DefaultClient *Client

func init() {
	DefaultClient = &Client{
		Key:        os.Getenv(API_KEY_ENV_VAR),
		HttpClient: http.DefaultClient,
	}
}

func (c *Client) parseListResponse(l []string, r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		l = append(l, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return l, nil
}

func (c *Client) List(l []string) ([]string, error) {
	url := urlForListRequest(nil, c.Key)
	req, err := http.NewRequest("POST", url.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	limitBody := io.LimitReader(resp.Body, 8196)
	l, err = c.parseListResponse(l, limitBody)
	if err != nil {
		return nil, err
	}

	return l, nil
}

func urlForListRequest(u *url.URL, key string) *url.URL {
	if u == nil {
		u = &url.URL{}
	}

	u.Scheme = API_SCHEME
	u.Host = API_HOST
	u.Path = API_LIST_PATH
	u.RawQuery = (url.Values{
		"client": []string{"api"},
		"apikey": []string{key},
		"appver": []string{CLIENT_VERSION},
		"pver":   []string{PROTOCOL_VERSION},
	}).Encode()

	return u
}
