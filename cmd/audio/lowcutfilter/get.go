package lowcutfilter

import (
	"encoding/json"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetLowCutFilterCmd represents the get audio low-cut-filter command
var GetLowCutFilterCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the audio low cut filter for a channel",
	RunE:  DoGetLowCutFilter,
}

func DoGetLowCutFilter(cmd *cobra.Command, args []string) error {
	c, err := client.New(viper.GetString("api.server"))
	if err != nil {
		return err
	}

	channelIndex, _ := cmd.Flags().GetInt("channel")
	filter, err := c.AudioGetLowCutFilter(channelIndex)
	if err != nil {
		return err
	}

	filterObj := client.AudioLowCutFilter{LowCutFilter: filter}
	out, err := json.MarshalIndent(filterObj, "", "  ")
	if err != nil {
		return err
	}
	cmd.Println(string(out))

	return nil
}

func init() {
	GetLowCutFilterCmd.Flags().IntP("channel", "c", 0, "The channel index to get the low cut filter for")
}
