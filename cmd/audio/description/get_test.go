package description_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/andrewhowdencom/bmctl/cmd/audio/description"
	"github.com/spf13/viper"
)

func TestGetDescriptionCmd(t *testing.T) {
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
		_ = json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	viper.Set("api.server", server.URL)

	b := new(bytes.Buffer)
	description.DescriptionCmd.SetOut(b)
	description.DescriptionCmd.SetArgs([]string{"get", "--channel", "0"})

	err := description.DescriptionCmd.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	output := b.String()
	if !strings.Contains(output, "Gain Min") || !strings.Contains(output, "-60") {
		t.Errorf("Output did not contain expected body. Output: %s", output)
	}
}
