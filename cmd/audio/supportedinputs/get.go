package supportedinputs

import (
	"fmt"
	"text/tabwriter"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetSupportedInputsCmd represents the get audio supported inputs command
var GetSupportedInputsCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the audio supported inputs for a channel",
	RunE:  DoGetSupportedInputs,
}

func DoGetSupportedInputs(cmd *cobra.Command, args []string) error {
	c, err := client.New(viper.GetString("api.server"))
	if err != nil {
		return err
	}

	channelIndex, _ := cmd.Flags().GetInt("channel")
	inputs, err := c.AudioGetSupportedInputs(channelIndex)
	if err != nil {
		return err
	}

	w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 0, 2, ' ', 0)
	_, _ = fmt.Fprintln(w, "INPUT\tAVAILABLE")
	for _, input := range inputs.SupportedInputs {
		_, _ = fmt.Fprintf(w, "%s\t%v\n", input.Schema.Input, input.Available)
	}
	_ = w.Flush()

	return nil
}

func init() {
	GetSupportedInputsCmd.Flags().IntP("channel", "c", 0, "The channel index to get the supported inputs for")
}
