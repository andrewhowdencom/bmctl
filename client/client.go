package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/moul/http2curl"
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

	http    *http.Client
	curlOut io.Writer
}

func New(addr string, opts ...Option) (*Client, error) {
	c := &Client{
		addr: addr,

		http: http.DefaultClient,
	}

	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

// WithCurlOutput configures the client to print an equivalent curl command to w
// before executing the request.
func WithCurlOutput(w io.Writer) Option {
	return func(c *Client) error {
		c.curlOut = w
		return nil
	}
}

// Do runs an arbitrary HTTP request to the server, filling the object (if possible) with the
func (c *Client) Do(in *http.Request, out interface{}) error {
	if c.curlOut != nil {
		cmd, err := http2curl.GetCurlCommand(in)
		if err != nil {
			return fmt.Errorf("%w: %s", ErrFailedHTTPRequestDo, err)
		}

		fmt.Fprintln(c.curlOut, cmd)
	}

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
	b, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrMarshalFailed, err)
	}

	req, err := http.NewRequest(method, c.addr+path, bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrFailedHTTPRequestCreation, err)
	}

	req.Header.Add("Content-Type", "application/json")

	return req, nil
}
