package supportedinputs_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/andrewhowdencom/bmctl/cmd/audio/supportedinputs"
	"github.com/spf13/viper"
)

func TestGetSupportedInputsCmd(t *testing.T) {
	expectedResponse := client.AudioSupportedInputs{
		SupportedInputs: []client.SupportedInput{
			{
				Schema:    client.AudioInput{Input: "Camera - Left"},
				Available: true,
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/control/api/v1/audio/channel/1/supportedInputs" {
			t.Errorf("Expected path /control/api/v1/audio/channel/1/supportedInputs, got %s", r.URL.Path)
		}
		_ = json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()

	viper.Set("api.server", server.URL)

	b := new(bytes.Buffer)
	supportedinputs.SupportedInputsCmd.SetOut(b)
	supportedinputs.SupportedInputsCmd.SetArgs([]string{"get", "--channel", "1"})

	err := supportedinputs.SupportedInputsCmd.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	output := b.String()
	if !strings.Contains(output, "Camera - Left") || !strings.Contains(output, "true") {
		t.Errorf("Output did not contain expected body. Output: %s", output)
	}
}
