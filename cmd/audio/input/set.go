package input

import (
	"errors"
	"fmt"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// SetInputCmd represents the set audio input command
var SetInputCmd = &cobra.Command{
	Use:   "set [input]",
	Short: "Set the audio input for a channel",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires an input argument")
		}
		return nil
	},
	RunE: DoSetInput,
}

func DoSetInput(cmd *cobra.Command, args []string) error {
	client, err := client.New(viper.GetString("api.server"))
	if err != nil {
		return err
	}

	all, _ := cmd.Flags().GetBool("all")
	if all {
		err = client.AudioSetAllInputs(args[0])
		if err != nil {
			return err
		}
		fmt.Println("Audio input set for all channels successfully")
		return nil
	}

	channelIndex, _ := cmd.Flags().GetInt("channel")
	err = client.AudioSetInput(channelIndex, args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Audio input for channel %d set successfully\n", channelIndex)

	return nil
}

func init() {
	SetInputCmd.Flags().IntP("channel", "c", 0, "The channel index to set the input for")
	SetInputCmd.Flags().Bool("all", false, "Set the input for all channels")
}