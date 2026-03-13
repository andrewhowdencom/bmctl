package supportedinputs

import (
	"fmt"
	"text/tabwriter"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var GetSupportedInputsCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the audio supported inputs for a channel",
	Long: `Retrieves a list of all supported audio input schemas mapped to their respective channel availability.
This command helps you discover exactly which input label strings are legally allowed to be passed into 'bmctl audio input set'.
The response will show hardware ports (like "XLR 1", "3.5mm Left", "AES/EBU") alongside whether they are active.`,
	Example: `  bmctl audio supported-inputs get --channel 1
  bmctl audio supported-inputs get -c 2`,
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
		_, _ = fmt.Fprintf(w, "%s\t%v\n", input.Input, input.Available)
	}
	_ = w.Flush()

	return nil
}

func init() {
	GetSupportedInputsCmd.Flags().IntP("channel", "c", 0, "The channel index to get the supported inputs for")
}
