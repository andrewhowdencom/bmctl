package available_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/andrewhowdencom/bmctl/cmd/audio/available"
	"github.com/spf13/viper"
)

func TestGetAvailableCmd(t *testing.T) {
	expectedResponse := client.AudioAvailable{Available: true}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/control/api/v1/audio/channel/1/available" {
			t.Errorf("Expected path /control/api/v1/audio/channel/1/available, got %s", r.URL.Path)
		}
		_ = json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	viper.Set("api.server", server.URL)

	b := new(bytes.Buffer)
	available.AvailableCmd.SetOut(b)
	available.AvailableCmd.SetArgs([]string{"get", "--channel", "1"})

	err := available.AvailableCmd.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	output := b.String()
	if !strings.Contains(output, "Available") || !strings.Contains(output, "true") {
		t.Errorf("Output did not contain expected body. Output: %s", output)
	}
}
