package client

import (
	"net/http"
)

// LensAutoFocusV1 performs auto focus
func (c *Client) LensAutoFocusV1() error {
	req, err := c.NewRequest(http.MethodPut, "/control/api/v1/lens/focus/doAutoFocus", nil)
	if err != nil {
		return err
	}

	return c.Do(req, nil)
}
