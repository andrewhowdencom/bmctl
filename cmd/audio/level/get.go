package level

import (
	"encoding/json"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetLevelCmd represents the get audio level command
var GetLevelCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the audio level for a channel",
	RunE:  DoGetLevel,
}

func DoGetLevel(cmd *cobra.Command, args []string) error {
	c, err := client.New(viper.GetString("api.server"))
	if err != nil {
		return err
	}

	channelIndex, _ := cmd.Flags().GetInt("channel")
	level, err := c.AudioGetLevel(channelIndex)
	if err != nil {
		return err
	}

	out, err := json.MarshalIndent(level, "", "  ")
	if err != nil {
		return err
	}
	cmd.Println(string(out))

	return nil
}

func init() {
	GetLevelCmd.Flags().IntP("channel", "c", 0, "The channel index to get the level for")
}
