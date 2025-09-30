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

	channelIndex, _ := cmd.Flags().GetInt("channel")
	err = client.AudioSetInput(channelIndex, args[0])
	if err != nil {
		return err
	}

	fmt.Println("Audio input set successfully")

	return nil
}

func init() {
	SetInputCmd.Flags().IntP("channel", "c", 0, "The channel index to set the input for")
}