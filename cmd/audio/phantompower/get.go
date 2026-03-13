package phantompower

import (
	"fmt"
	"text/tabwriter"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var GetPhantomPowerCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the audio phantom power for a channel",
	Long: `Retrieves the current state of 48V Phantom Power (true/false) on the specified audio channel.
Phantom power is required to drive active microphones like condensers, which do not have an internal battery.
It is typically applied to XLR connections. Ensure Phantom Power is OFF before plugging in ribbon microphones as it may damage them.`,
	Example: `  bmctl audio phantom-power get --channel 1
  bmctl audio phantom-power get -c 2`,
	RunE:  DoGetPhantomPower,
}

func DoGetPhantomPower(cmd *cobra.Command, args []string) error {
	c, err := client.New(viper.GetString("api.server"))
	if err != nil {
		return err
	}

	channelIndex, _ := cmd.Flags().GetInt("channel")
	power, err := c.AudioGetPhantomPower(channelIndex)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
	_, _ = fmt.Fprintln(w, "PROPERTY\tVALUE")
	_, _ = fmt.Fprintf(w, "Phantom Power\t%v\n", power)
	_ = w.Flush()

	return nil
}

func init() {
	GetPhantomPowerCmd.Flags().IntP("channel", "c", 0, "The channel index to get the phantom power for")
}
