package input

import (
	"fmt"
	"text/tabwriter"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetInputCmd represents the get audio input command
var GetInputCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the audio input for a channel",
	Long: `Retrieves the name of the active audio input source currently routed to the specified channel.
Common outputs include "Camera - Left", "Camera - Right", "XLR 1", "XLR 2", or "3.5mm Left".`,
	Example: `  bmctl audio input get --channel 1
  bmctl audio input get -c 2`,
	RunE:  DoGetInput,
}

func DoGetInput(cmd *cobra.Command, args []string) error {
	client, err := client.New(viper.GetString("api.server"))
	if err != nil {
		return err
	}

	channelIndex, _ := cmd.Flags().GetInt("channel")
	input, err := client.AudioGetInput(channelIndex)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
	_, _ = fmt.Fprintln(w, "PROPERTY\tVALUE")
	_, _ = fmt.Fprintf(w, "Input\t%s\n", input)
	_ = w.Flush()

	return nil
}

func init() {
	GetInputCmd.Flags().IntP("channel", "c", 0, "The channel index to get the input for")
}