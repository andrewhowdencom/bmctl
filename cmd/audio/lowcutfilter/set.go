package lowcutfilter

import (
	"fmt"
	"strconv"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var SetLowCutFilterCmd = &cobra.Command{
	Use:   "set [true|false]",
	Short: "Set the audio low cut filter for a channel",
	Long: `Enables or disables the low cut filter for the specified
audio channel. When set to true, low-frequency rumble (like wind or
handling noise) is filtered out before recording. Use the '--all'
flag to bulk apply this filter boolean to all active channels.`,
	Example: `  bmctl audio low-cut-filter set true --channel 1
  bmctl audio low-cut-filter set false -c 2
  bmctl audio low-cut-filter set true --all`,
	Args:  cobra.ExactArgs(1),
	RunE:  DoSetLowCutFilter,
}

func DoSetLowCutFilter(cmd *cobra.Command, args []string) error {
	c, err := client.New(viper.GetString("api.server"))
	if err != nil {
		return err
	}

	filter, err := strconv.ParseBool(args[0])
	if err != nil {
		return fmt.Errorf("invalid boolean argument: %w", err)
	}

	all, _ := cmd.Flags().GetBool("all")
	if all {
		err = c.AudioSetAllLowCutFilter(filter)
		if err != nil {
			return err
		}
		cmd.Println("Audio low cut filter set for all channels successfully")
		return nil
	}

	channelIndex, _ := cmd.Flags().GetInt("channel")
	err = c.AudioSetLowCutFilter(channelIndex, filter)
	if err != nil {
		return err
	}

	cmd.Printf("Audio low cut filter for channel %d set successfully\n", channelIndex)
	return nil
}

func init() {
	SetLowCutFilterCmd.Flags().IntP("channel", "c", 0, "The channel index to set the low cut filter for")
	SetLowCutFilterCmd.Flags().Bool("all", false, "Set the low cut filter for all channels")
}
