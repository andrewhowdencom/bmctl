package padding

import (
	"fmt"
	"text/tabwriter"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var GetPaddingCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the audio padding for a channel",
	Long: `Retrieves the current state of attenuation padding (true/false)
for the specified audio channel. Audio padding attenuates (lowers)
the incoming signal level before the pre-amp. It is extremely useful
to prevent digital clipping when recording very loud sound sources
via line or mic inputs.`,
	Example: `  bmctl audio padding get --channel 1
  bmctl audio padding get -c 2`,
	RunE:  DoGetPadding,
}

func DoGetPadding(cmd *cobra.Command, args []string) error {
	c, err := client.New(viper.GetString("api.server"))
	if err != nil {
		return err
	}

	channelIndex, _ := cmd.Flags().GetInt("channel")
	padding, err := c.AudioGetPadding(channelIndex)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
	_, _ = fmt.Fprintln(w, "PROPERTY\tVALUE")
	_, _ = fmt.Fprintf(w, "Padding\t%v\n", padding)
	_ = w.Flush()

	return nil
}

func init() {
	GetPaddingCmd.Flags().IntP("channel", "c", 0, "The channel index to get the padding for")
}
