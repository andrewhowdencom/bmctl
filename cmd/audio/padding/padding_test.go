package padding_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/andrewhowdencom/bmctl/cmd/audio/padding"
	"github.com/spf13/viper"
)

func TestGetPaddingCmd(t *testing.T) {
	expectedResponse := client.AudioPadding{Padding: true}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/control/api/v1/audio/channel/1/padding" {
			t.Errorf("Expected path /control/api/v1/audio/channel/1/padding, got %s", r.URL.Path)
		}
		_ = json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	viper.Set("api.server", server.URL)

	b := new(bytes.Buffer)
	padding.PaddingCmd.SetOut(b)
	padding.PaddingCmd.SetArgs([]string{"get", "--channel", "1"})

	err := padding.PaddingCmd.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	output := b.String()
	if !strings.Contains(output, "Padding") || !strings.Contains(output, "true") {
		t.Errorf("Output did not contain expected body. Output: %s", output)
	}
}

func TestSetPaddingCmd(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/control/api/v1/audio/channel/0/padding" {
			t.Errorf("Expected path /control/api/v1/audio/channel/0/padding, got %s", r.URL.Path)
		}
		
		var reqBody client.AudioPadding
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			t.Errorf("Failed to decode request body: %v", err)
		}

		if reqBody.Padding {
			t.Errorf("Expected Padding to be false, got %v", reqBody.Padding)
		}
		
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	viper.Set("api.server", server.URL)

	b := new(bytes.Buffer)
	padding.PaddingCmd.SetOut(b)
	padding.PaddingCmd.SetArgs([]string{"set", "--channel", "0", "false"})

	err := padding.PaddingCmd.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}
}
