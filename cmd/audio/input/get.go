package input

import (
	"fmt"
	"os"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetInputCmd represents the get audio input command
var GetInputCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the audio input for a channel",
	RunE:  DoGetInput,
}

func DoGetInput(cmd *cobra.Command, args []string) error {
	opts := []client.Option{}
	if viper.GetBool("curl") {
		opts = append(opts, client.WithCurlOutput(os.Stderr))
	}

	client, err := client.New(viper.GetString("api.server"), opts...)
	if err != nil {
		return err
	}

	channelIndex, _ := cmd.Flags().GetInt("channel")
	input, err := client.AudioGetInput(channelIndex)
	if err != nil {
		return err
	}

	fmt.Println(input)

	return nil
}

func init() {
	GetInputCmd.Flags().IntP("channel", "c", 0, "The channel index to get the input for")
}