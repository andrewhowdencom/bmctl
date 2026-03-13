package padding

import (
	"fmt"
	"strconv"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// SetPaddingCmd represents the set audio padding command
var SetPaddingCmd = &cobra.Command{
	Use:   "set [true|false]",
	Short: "Set the audio padding for a channel",
	Args:  cobra.ExactArgs(1),
	RunE:  DoSetPadding,
}

func DoSetPadding(cmd *cobra.Command, args []string) error {
	c, err := client.New(viper.GetString("api.server"))
	if err != nil {
		return err
	}

	padding, err := strconv.ParseBool(args[0])
	if err != nil {
		return fmt.Errorf("invalid boolean argument: %w", err)
	}

	all, _ := cmd.Flags().GetBool("all")
	if all {
		err = c.AudioSetAllPadding(padding)
		if err != nil {
			return err
		}
		cmd.Println("Audio padding set for all channels successfully")
		return nil
	}

	channelIndex, _ := cmd.Flags().GetInt("channel")
	err = c.AudioSetPadding(channelIndex, padding)
	if err != nil {
		return err
	}

	cmd.Printf("Audio padding for channel %d set successfully\n", channelIndex)
	return nil
}

func init() {
	SetPaddingCmd.Flags().IntP("channel", "c", 0, "The channel index to set the padding for")
	SetPaddingCmd.Flags().Bool("all", false, "Set the padding for all channels")
}
