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

func (c *Client) AudioSetAllLowCutFilter(filter bool) error {
	for i := 0; i < 2; i++ {
		if err := c.AudioSetLowCutFilter(i, filter); err != nil {
			return fmt.Errorf("failed to set low cut filter for channel %d: %w", i, err)
		}
	}
	return nil
}

func (c *Client) AudioSetAllPadding(padding bool) error {
	for i := 0; i < 2; i++ {
		if err := c.AudioSetPadding(i, padding); err != nil {
			return fmt.Errorf("failed to set padding for channel %d: %w", i, err)
		}
	}
	return nil
}

func (c *Client) AudioSetAllPhantomPower(power bool) error {
	for i := 0; i < 2; i++ {
		if err := c.AudioSetPhantomPower(i, power); err != nil {
			return fmt.Errorf("failed to set phantom power for channel %d: %w", i, err)
		}
	}
	return nil
}

func (c *Client) AudioSetAllLevels(level AudioLevel) error {
	for i := 0; i < 2; i++ {
		if err := c.AudioSetLevel(i, level); err != nil {
			return fmt.Errorf("failed to set level for channel %d: %w", i, err)
		}
	}
	return nil
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
type AudioInputDescription struct {
GainRange    GainRange    `json:"gainRange"`
Capabilities Capabilities `json:"capabilities"`
}

type GainRange struct {
Min float64 `json:"Min"`
Max float64 `json:"Max"`
}

type Capabilities struct {
PhantomPower bool         `json:"PhantomPower"`
LowCutFilter bool         `json:"LowCutFilter"`
Padding      PaddingDesc  `json:"Padding"`
}

type PaddingDesc struct {
Available bool    `json:"available"`
Forced    bool    `json:"forced"`
Value     float64 `json:"value"`
}

type AudioSupportedInputs struct {
SupportedInputs []SupportedInput `json:"supportedInputs"`
}

type SupportedInput struct {
Schema    AudioInput `json:"schema"`
Available bool       `json:"available"`
}

type AudioLevel struct {
Gain       *float64 `json:"gain,omitempty"`
Normalised *float64 `json:"normalised,omitempty"`
}

type AudioPhantomPower struct {
PhantomPower bool `json:"phantomPower"`
}

type AudioPadding struct {
Padding bool `json:"padding"`
}

type AudioLowCutFilter struct {
LowCutFilter bool `json:"lowCutFilter"`
}

type AudioAvailable struct {
Available bool `json:"available"`
}

func (c *Client) AudioGetDescription(channelIndex int) (*AudioInputDescription, error) {
path := fmt.Sprintf("/control/api/v1/audio/channel/%d/input/description", channelIndex)
req, err := c.NewRequest(http.MethodGet, path, nil)
if err != nil {
return nil, err
}

out := AudioInputDescription{}
if err := c.Do(req, &out); err != nil {
return nil, err
}

return &out, nil
}

func (c *Client) AudioGetSupportedInputs(channelIndex int) (*AudioSupportedInputs, error) {
path := fmt.Sprintf("/control/api/v1/audio/channel/%d/supportedInputs", channelIndex)
req, err := c.NewRequest(http.MethodGet, path, nil)
if err != nil {
return nil, err
}

out := AudioSupportedInputs{}
if err := c.Do(req, &out); err != nil {
return nil, err
}

return &out, nil
}

func (c *Client) AudioGetLevel(channelIndex int) (*AudioLevel, error) {
path := fmt.Sprintf("/control/api/v1/audio/channel/%d/level", channelIndex)
req, err := c.NewRequest(http.MethodGet, path, nil)
if err != nil {
return nil, err
}

out := AudioLevel{}
if err := c.Do(req, &out); err != nil {
return nil, err
}

return &out, nil
}

func (c *Client) AudioSetLevel(channelIndex int, level AudioLevel) error {
path := fmt.Sprintf("/control/api/v1/audio/channel/%d/level", channelIndex)
req, err := c.NewRequest(http.MethodPut, path, &level)
if err != nil {
return err
}

return c.Do(req, nil)
}

func (c *Client) AudioGetPhantomPower(channelIndex int) (bool, error) {
path := fmt.Sprintf("/control/api/v1/audio/channel/%d/phantomPower", channelIndex)
req, err := c.NewRequest(http.MethodGet, path, nil)
if err != nil {
return false, err
}

out := AudioPhantomPower{}
if err := c.Do(req, &out); err != nil {
return false, err
}

return out.PhantomPower, nil
}

func (c *Client) AudioSetPhantomPower(channelIndex int, power bool) error {
path := fmt.Sprintf("/control/api/v1/audio/channel/%d/phantomPower", channelIndex)
in := AudioPhantomPower{PhantomPower: power}
req, err := c.NewRequest(http.MethodPut, path, &in)
if err != nil {
return err
}

return c.Do(req, nil)
}

func (c *Client) AudioGetPadding(channelIndex int) (bool, error) {
path := fmt.Sprintf("/control/api/v1/audio/channel/%d/padding", channelIndex)
req, err := c.NewRequest(http.MethodGet, path, nil)
if err != nil {
return false, err
}

out := AudioPadding{}
if err := c.Do(req, &out); err != nil {
return false, err
}

return out.Padding, nil
}

func (c *Client) AudioSetPadding(channelIndex int, padding bool) error {
path := fmt.Sprintf("/control/api/v1/audio/channel/%d/padding", channelIndex)
in := AudioPadding{Padding: padding}
req, err := c.NewRequest(http.MethodPut, path, &in)
if err != nil {
return err
}

return c.Do(req, nil)
}

func (c *Client) AudioGetLowCutFilter(channelIndex int) (bool, error) {
path := fmt.Sprintf("/control/api/v1/audio/channel/%d/lowCutFilter", channelIndex)
req, err := c.NewRequest(http.MethodGet, path, nil)
if err != nil {
return false, err
}

out := AudioLowCutFilter{}
if err := c.Do(req, &out); err != nil {
return false, err
}

return out.LowCutFilter, nil
}

func (c *Client) AudioSetLowCutFilter(channelIndex int, filter bool) error {
path := fmt.Sprintf("/control/api/v1/audio/channel/%d/lowCutFilter", channelIndex)
in := AudioLowCutFilter{LowCutFilter: filter}
req, err := c.NewRequest(http.MethodPut, path, &in)
if err != nil {
return err
}

return c.Do(req, nil)
}

func (c *Client) AudioGetAvailable(channelIndex int) (bool, error) {
path := fmt.Sprintf("/control/api/v1/audio/channel/%d/available", channelIndex)
req, err := c.NewRequest(http.MethodGet, path, nil)
if err != nil {
return false, err
}

out := AudioAvailable{}
if err := c.Do(req, &out); err != nil {
return false, err
}

return out.Available, nil
}
