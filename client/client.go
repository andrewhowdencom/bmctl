package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrFailedHTTPRequestCreation   = errors.New("unable to create http request")
	ErrFailedHTTPRequestDo         = errors.New("unable to do HTTP request")
	ErrUnsupportedHTTPResponseCode = errors.New("unsupported HTTP response code")

	ErrMarshalFailed   = errors.New("cannot marshal request")
	ErrUnmarshalFailed = errors.New("cannot unmarshal response")
)

type Option func(c *Client) error

type Client struct {
	// addr is the address of the client. For example, "http://192.168.1.1"
	addr string

	http *http.Client
}

func New(addr string, _ ...Option) (*Client, error) {
	return &Client{
		addr: addr,

		http: http.DefaultClient,
	}, nil
}

// Do runs an arbitrary HTTP request to the server, filling the object (if possible) with the
func (c *Client) Do(in *http.Request, out interface{}) error {
	resp, err := c.http.Do(in)

	if err != nil {
		return fmt.Errorf("%w: %s", ErrFailedHTTPRequestDo, err)
	}

	switch resp.StatusCode {
	case http.StatusOK:
		// The good case, continue
	case http.StatusNoContent:
		// If there is no content, we do not need to unserialize anything and can return early.
		return nil
	default:
		return fmt.Errorf("%w: %d", ErrUnsupportedHTTPResponseCode, resp.StatusCode)
	}

	d := json.NewDecoder(resp.Body)
	if err := d.Decode(out); err != nil {
		return fmt.Errorf("%w: %s", ErrUnmarshalFailed, err)
	}

	return nil
}

// NewRequest takes the input request object, and marshals it into the expected HTTP request, complete with headers.
func (c *Client) NewRequest(method string, path string, in interface{}) (*http.Request, error) {
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(in); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrMarshalFailed, err)
	}

	req, err := http.NewRequest(method, c.addr+path, buf)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrFailedHTTPRequestCreation, err)
	}

	req.Header.Add("Content-Type", "application/json")

	return req, nil
}
