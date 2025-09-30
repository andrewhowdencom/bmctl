package client

import (
	"fmt"
	"net/http"
)

// AudioInput defines the structure for an audio input request/response
type AudioInput struct {
	Input string `json:"input"`
}

// AudioGetInput gets the audio input for a specified channel
func (c *Client) AudioGetInput(channelIndex int) (string, error) {
	path := fmt.Sprintf("/control/api/v1/audio/channel/%d/input", channelIndex)
	req, err := c.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return "", err
	}

	out := AudioInput{}
	if err := c.Do(req, &out); err != nil {
		return "", err
	}

	return out.Input, nil
}

// AudioSetInput sets the audio input for a specified channel
func (c *Client) AudioSetInput(channelIndex int, input string) error {
	path := fmt.Sprintf("/control/api/v1/audio/channel/%d/input", channelIndex)
	in := AudioInput{
		Input: input,
	}

	req, err := c.NewRequest(http.MethodPut, path, &in)
	if err != nil {
		return err
	}

	return c.Do(req, nil)
}

// AudioSetAllInputs sets the audio input for all channels (0 and 1)
func (c *Client) AudioSetAllInputs(input string) error {
	// Assuming 2 channels based on user feedback ("left and right")
	for i := 0; i < 2; i++ {
		err := c.AudioSetInput(i, input)
		if err != nil {
			return fmt.Errorf("failed to set input for channel %d: %w", i, err)
		}
	}
	return nil
}