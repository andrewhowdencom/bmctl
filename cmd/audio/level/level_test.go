package level_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/andrewhowdencom/bmctl/cmd/audio/level"
	"github.com/spf13/viper"
)

func TestGetLevelCmd(t *testing.T) {
	gain := float64(4.5)
	expectedResponse := client.AudioLevel{
		Gain: &gain,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/control/api/v1/audio/channel/1/level" {
			t.Errorf("Expected path /control/api/v1/audio/channel/1/level, got %s", r.URL.Path)
		}
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	viper.Set("api.server", server.URL)

	b := new(bytes.Buffer)
	level.LevelCmd.SetOut(b)
	level.LevelCmd.SetArgs([]string{"get", "--channel", "1"})

	err := level.LevelCmd.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	output := b.String()
	if !strings.Contains(output, "gain") || !strings.Contains(output, "4.5") {
		t.Errorf("Output did not contain expected body. Output: %s", output)
	}
}

func TestSetLevelCmd_Gain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/control/api/v1/audio/channel/0/level" {
			t.Errorf("Expected path /control/api/v1/audio/channel/0/level, got %s", r.URL.Path)
		}
		
		var reqBody client.AudioLevel
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			t.Errorf("Failed to decode request body: %v", err)
		}

		if reqBody.Gain == nil || *reqBody.Gain != -5.5 {
			t.Errorf("Expected Gain -5.5, got %v", reqBody.Gain)
		}
		
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	viper.Set("api.server", server.URL)

	b := new(bytes.Buffer)
	level.LevelCmd.SetOut(b)
	level.LevelCmd.SetArgs([]string{"set", "--channel", "0", "--gain", "-5.5"})

	err := level.LevelCmd.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}
}

func TestSetLevelCmd_Normalised(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/control/api/v1/audio/channel/2/level" {
			t.Errorf("Expected path /control/api/v1/audio/channel/2/level, got %s", r.URL.Path)
		}
		
		var reqBody client.AudioLevel
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			t.Errorf("Failed to decode request body: %v", err)
		}

		if reqBody.Normalised == nil || *reqBody.Normalised != 0.75 {
			t.Errorf("Expected Normalised 0.75, got %v", reqBody.Normalised)
		}
		if reqBody.Gain != nil {
			t.Errorf("Expected Gain nil, got %v", reqBody.Gain)
		}
		
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	viper.Set("api.server", server.URL)

	b := new(bytes.Buffer)
	level.LevelCmd.SetOut(b)
	level.LevelCmd.SetArgs([]string{"set", "--channel", "2", "--normalised", "0.75"})

	err := level.LevelCmd.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}
}

func TestSetLevelCmd_NoArgs(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	viper.Set("api.server", server.URL)

	b := new(bytes.Buffer)
	level.LevelCmd.SetOut(b)
	level.LevelCmd.SetArgs([]string{"set", "--channel", "0"})

	err := level.LevelCmd.Execute()
	if err == nil {
		t.Fatalf("Expected an error when executing without gain or normalised args")
	}
}
