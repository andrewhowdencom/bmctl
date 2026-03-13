package phantompower_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/andrewhowdencom/bmctl/cmd/audio/phantompower"
	"github.com/spf13/viper"
)

func TestGetPhantomPowerCmd(t *testing.T) {
	expectedResponse := client.AudioPhantomPower{PhantomPower: true}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/control/api/v1/audio/channel/1/phantomPower" {
			t.Errorf("Expected path /control/api/v1/audio/channel/1/phantomPower, got %s", r.URL.Path)
		}
		json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	viper.Set("api.server", server.URL)

	b := new(bytes.Buffer)
	phantompower.PhantomPowerCmd.SetOut(b)
	phantompower.PhantomPowerCmd.SetArgs([]string{"get", "--channel", "1"})

	err := phantompower.PhantomPowerCmd.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	output := b.String()
	if !strings.Contains(output, "phantomPower") || !strings.Contains(output, "true") {
		t.Errorf("Output did not contain expected body. Output: %s", output)
	}
}

func TestSetPhantomPowerCmd(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/control/api/v1/audio/channel/0/phantomPower" {
			t.Errorf("Expected path /control/api/v1/audio/channel/0/phantomPower, got %s", r.URL.Path)
		}
		
		var reqBody client.AudioPhantomPower
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			t.Errorf("Failed to decode request body: %v", err)
		}

		if !reqBody.PhantomPower {
			t.Errorf("Expected PhantomPower to be true, got %v", reqBody.PhantomPower)
		}
		
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	viper.Set("api.server", server.URL)

	b := new(bytes.Buffer)
	phantompower.PhantomPowerCmd.SetOut(b)
	phantompower.PhantomPowerCmd.SetArgs([]string{"set", "--channel", "0", "true"})

	err := phantompower.PhantomPowerCmd.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}
}
