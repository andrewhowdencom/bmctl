package phantompower

import (
	"fmt"
	"text/tabwriter"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetPhantomPowerCmd represents the get audio phantom-power command
var GetPhantomPowerCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the audio phantom power for a channel",
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
