package client

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrOutOfRange = errors.New("out of acceptable range")
)

var (
	ValidISO = []int{
		100, 125, 160, 200, 250, 320, 400, 500, 640, 800, 1000, 1250, 1600, 2000,
		2500, 3200, 4000, 5000, 6400, 8000, 10000, 12800, 16000, 20000, 2560,
	}

	ValidISOMap = func() map[int]struct{} {
		m := make(map[int]struct{})
		for _, v := range ValidISO {
			m[v] = struct{}{}
		}

		return m
	}()
)

type ISO struct {
	Value int `json:"iso"`
}

const (
	VideoISOMax, VideoISOMin = 2147483647, 0
)

// VideoSetISO sets the ISO of the video
func (c *Client) VideoSetISO(iso int) error {
	body := &ISO{
		Value: iso,
	}

	if iso < VideoISOMin || iso > VideoISOMax {
		return fmt.Errorf("%w: %d-%d (%d)", ErrOutOfRange, VideoISOMin, VideoISOMax, iso)
	}

	if _, ok := ValidISOMap[iso]; !ok {
		return fmt.Errorf("%w: %s (%v)", ErrOutOfRange, "can only use ISOs in range", ValidISO)
	}

	req, err := c.NewRequest(http.MethodPut, "/control/api/v1/video/iso", body)
	if err != nil {
		return err
	}

	return c.Do(req, nil)
}

// VideoSetISO sets the ISO of the video
func (c *Client) VideoGetISO() (int, error) {
	req, err := c.NewRequest(http.MethodGet, "/control/api/v1/video/iso", nil)
	if err != nil {
		return 0, err
	}

	resp := &ISO{}

	if err := c.Do(req, resp); err != nil {
		return 0, err
	}

	return resp.Value, nil
}
