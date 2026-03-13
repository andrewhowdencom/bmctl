package padding

import (
	"fmt"
	"text/tabwriter"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetPaddingCmd represents the get audio padding command
var GetPaddingCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the audio padding for a channel",
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
