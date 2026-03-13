package description

import (
	"fmt"
	"text/tabwriter"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetDescriptionCmd represents the get audio description command
var GetDescriptionCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the audio input description for a channel",
	Long: `Retrieves a detailed description of the audio input assigned to the specified channel.
This includes hardware capabilities such as the physical gain range (e.g., -60dB to 12dB),
and whether the input supports Phantom Power (48V), Low Cut Filters, or Attenuation Padding (-10dB/-20dB).`,
	Example: `  bmctl audio description get --channel 1
  bmctl audio description get -c 2`,
	RunE:  DoGetDescription,
}

func DoGetDescription(cmd *cobra.Command, args []string) error {
	c, err := client.New(viper.GetString("api.server"))
	if err != nil {
		return err
	}

	channelIndex, _ := cmd.Flags().GetInt("channel")
	desc, err := c.AudioGetDescription(channelIndex)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
	_, _ = fmt.Fprintln(w, "PROPERTY\tVALUE")
	_, _ = fmt.Fprintf(w, "Gain Min\t%v\n", desc.GainRange.Min)
	_, _ = fmt.Fprintf(w, "Gain Max\t%v\n", desc.GainRange.Max)
	_, _ = fmt.Fprintf(w, "Phantom Power\t%v\n", desc.Capabilities.PhantomPower)
	_, _ = fmt.Fprintf(w, "Low Cut Filter\t%v\n", desc.Capabilities.LowCutFilter)
	_, _ = fmt.Fprintf(w, "Padding Available\t%v\n", desc.Capabilities.Padding.Available)
	_, _ = fmt.Fprintf(w, "Padding Forced\t%v\n", desc.Capabilities.Padding.Forced)
	_, _ = fmt.Fprintf(w, "Padding Value\t%v\n", desc.Capabilities.Padding.Value)
	_ = w.Flush()

	return nil
}

func init() {
	GetDescriptionCmd.Flags().IntP("channel", "c", 0, "The channel index to get the description for")
}
