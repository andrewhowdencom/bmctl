package client_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andrewhowdencom/bmctl/client"
)

func TestAudioGetDescription(t *testing.T) {
	expectedResponse := client.AudioInputDescription{
		GainRange: client.GainRange{Min: -60, Max: 12},
		Capabilities: client.Capabilities{
			PhantomPower: true,
			LowCutFilter: true,
			Padding: client.PaddingDesc{
				Available: true,
				Forced:    false,
				Value:     -10.0,
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/control/api/v1/audio/channel/0/input/description" {
			t.Errorf("Expected path /control/api/v1/audio/channel/0/input/description, got %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	c, err := client.New(server.URL)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	res, err := c.AudioGetDescription(0)
	if err != nil {
		t.Fatalf("AudioGetDescription failed: %v", err)
	}

	if res.GainRange.Min != expectedResponse.GainRange.Min {
		t.Errorf("Expected GainRange.Min %v, got %v", expectedResponse.GainRange.Min, res.GainRange.Min)
	}
	if res.Capabilities.PhantomPower != expectedResponse.Capabilities.PhantomPower {
		t.Errorf("Expected Capabilities.PhantomPower %v, got %v", expectedResponse.Capabilities.PhantomPower, res.Capabilities.PhantomPower)
	}
}

func TestAudioGetLevel(t *testing.T) {
	gain := float64(5.0)
	var expectedResponse = client.AudioLevel{
		Gain: &gain,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/control/api/v1/audio/channel/1/level" {
			t.Errorf("Expected path /control/api/v1/audio/channel/1/level, got %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	c, err := client.New(server.URL)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	res, err := c.AudioGetLevel(1)
	if err != nil {
		t.Fatalf("AudioGetLevel failed: %v", err)
	}

	if res.Gain == nil || *res.Gain != 5.0 {
		t.Errorf("Expected Gain 5.0, got %v", res.Gain)
	}
	if res.Normalised != nil {
		t.Errorf("Expected Normalised nil, got %v", res.Normalised)
	}
}

func TestAudioSetLevel(t *testing.T) {
	gain := float64(-12.5)
	levelToSet := client.AudioLevel{
		Gain: &gain,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/control/api/v1/audio/channel/0/level" {
			t.Errorf("Expected path /control/api/v1/audio/channel/0/level, got %s", r.URL.Path)
		}
		if r.Method != http.MethodPut {
			t.Errorf("Expected method PUT, got %s", r.Method)
		}

		var reqBody client.AudioLevel
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			t.Errorf("Failed to decode request body: %v", err)
		}

		if reqBody.Gain == nil || *reqBody.Gain != -12.5 {
			t.Errorf("Expected request body Gain to be -12.5, got %v", reqBody.Gain)
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	c, err := client.New(server.URL)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	err = c.AudioSetLevel(0, levelToSet)
	if err != nil {
		t.Fatalf("AudioSetLevel failed: %v", err)
	}
}

func TestAudioGetSetPhantomPower(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/control/api/v1/audio/channel/0/phantomPower" {
			t.Errorf("Expected path ..., got %s", r.URL.Path)
		}

		if r.Method == http.MethodGet {
			json.NewEncoder(w).Encode(client.AudioPhantomPower{PhantomPower: true})
		} else if r.Method == http.MethodPut {
			var reqBody client.AudioPhantomPower
			if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
				t.Errorf("Failed to decode request body: %v", err)
			}
			if !reqBody.PhantomPower {
				t.Errorf("Expected request body PhantomPower to be true")
			}
			w.WriteHeader(http.StatusOK)
		}
	}))
	defer server.Close()

	c, err := client.New(server.URL)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	res, err := c.AudioGetPhantomPower(0)
	if err != nil {
		t.Fatalf("AudioGetPhantomPower failed: %v", err)
	}
	if !res {
		t.Errorf("Expected true, got %v", res)
	}

	err = c.AudioSetPhantomPower(0, true)
	if err != nil {
		t.Fatalf("AudioSetPhantomPower failed: %v", err)
	}
}
