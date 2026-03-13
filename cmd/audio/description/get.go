package description

import (
	"encoding/json"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetDescriptionCmd represents the get audio description command
var GetDescriptionCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the audio input description for a channel",
	RunE:  DoGetDescription,
}

func DoGetDescription(cmd *cobra.Command, args []string) error {
	c, err := client.New(viper.GetString("api.server"))
	if err != nil {
		return err
	}

	channelIndex, _ := cmd.Flags().GetInt("channel")
	desc, err := c.AudioGetDescription(channelIndex)
	if err != nil {
		return err
	}

	out, err := json.MarshalIndent(desc, "", "  ")
	if err != nil {
		return err
	}
	cmd.Println(string(out))

	return nil
}

func init() {
	GetDescriptionCmd.Flags().IntP("channel", "c", 0, "The channel index to get the description for")
}
