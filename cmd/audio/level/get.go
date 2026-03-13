package level

import (
	"fmt"
	"text/tabwriter"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetLevelCmd represents the get audio level command
var GetLevelCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the audio level for a channel",
	Long: `Retrieves the current audio gain level for the specified channel.
Audio levels can be returned as either absolute physical 'Gain' values
in decibels (dB), or alternatively as a 'Normalised' floating point
value (0.0 to 1.0) depending on the camera hardware's precision schema.`,
	Example: `  bmctl audio level get --channel 1
  bmctl audio level get -c 2`,
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

	w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
	_, _ = fmt.Fprintln(w, "PROPERTY\tVALUE")
	if level.Gain != nil {
		_, _ = fmt.Fprintf(w, "Gain\t%v\n", *level.Gain)
	} else {
		_, _ = fmt.Fprintln(w, "Gain\t<nil>")
	}
	if level.Normalised != nil {
		_, _ = fmt.Fprintf(w, "Normalised\t%v\n", *level.Normalised)
	} else {
		_, _ = fmt.Fprintln(w, "Normalised\t<nil>")
	}
	_ = w.Flush()

	return nil
}

func init() {
	GetLevelCmd.Flags().IntP("channel", "c", 0, "The channel index to get the level for")
}
