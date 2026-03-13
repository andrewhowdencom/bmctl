package lowcutfilter

import (
	"fmt"
	"text/tabwriter"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var GetLowCutFilterCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the audio low cut filter for a channel",
	Long: `Retrieves the current state of the low cut filter (true/false)
for the specified audio channel. A low cut filter is typically used
on microphones to reduce low-frequency rumble such as wind noise,
handling noise, or vocal pops.`,
	Example: `  bmctl audio low-cut-filter get --channel 1
  bmctl audio low-cut-filter get -c 2`,
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

	w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
	_, _ = fmt.Fprintln(w, "PROPERTY\tVALUE")
	_, _ = fmt.Fprintf(w, "Low Cut Filter\t%v\n", filter)
	_ = w.Flush()

	return nil
}

func init() {
	GetLowCutFilterCmd.Flags().IntP("channel", "c", 0, "The channel index to get the low cut filter for")
}
