package supportedinputs

import (
	"encoding/json"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetSupportedInputsCmd represents the get audio supported inputs command
var GetSupportedInputsCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the audio supported inputs for a channel",
	RunE:  DoGetSupportedInputs,
}

func DoGetSupportedInputs(cmd *cobra.Command, args []string) error {
	c, err := client.New(viper.GetString("api.server"))
	if err != nil {
		return err
	}

	channelIndex, _ := cmd.Flags().GetInt("channel")
	inputs, err := c.AudioGetSupportedInputs(channelIndex)
	if err != nil {
		return err
	}

	out, err := json.MarshalIndent(inputs, "", "  ")
	if err != nil {
		return err
	}
	cmd.Println(string(out))

	return nil
}

func init() {
	GetSupportedInputsCmd.Flags().IntP("channel", "c", 0, "The channel index to get the supported inputs for")
}
