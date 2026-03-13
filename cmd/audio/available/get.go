package available

import (
	"encoding/json"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetAvailableCmd represents the get audio available command
var GetAvailableCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the audio input's current availability for the selected channel",
	RunE:  DoGetAvailable,
}

func DoGetAvailable(cmd *cobra.Command, args []string) error {
	c, err := client.New(viper.GetString("api.server"))
	if err != nil {
		return err
	}

	channelIndex, _ := cmd.Flags().GetInt("channel")
	available, err := c.AudioGetAvailable(channelIndex)
	if err != nil {
		return err
	}

	availableObj := client.AudioAvailable{Available: available}
	out, err := json.MarshalIndent(availableObj, "", "  ")
	if err != nil {
		return err
	}
	cmd.Println(string(out))

	return nil
}

func init() {
	GetAvailableCmd.Flags().IntP("channel", "c", 0, "The channel index to get the availability for")
}
