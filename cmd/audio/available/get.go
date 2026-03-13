package available

import (
	"fmt"
	"text/tabwriter"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetAvailableCmd represents the get audio available command
var GetAvailableCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the audio input's current availability for the selected channel",
	Long: `Returns whether the audio input for the specified channel is
currently available. Some audio inputs (e.g. physical ports like XLR
or 3.5mm) might not be available if they are disconnected or
unsupported by the current recording mode.`,
	Example: `  bmctl audio available get --channel 1
  bmctl audio available get -c 2`,
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

	w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
	_, _ = fmt.Fprintln(w, "PROPERTY\tVALUE")
	_, _ = fmt.Fprintf(w, "Available\t%v\n", available)
	_ = w.Flush()

	return nil
}

func init() {
	GetAvailableCmd.Flags().IntP("channel", "c", 0, "The channel index to get the availability for")
}
