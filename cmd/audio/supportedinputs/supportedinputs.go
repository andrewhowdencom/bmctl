package supportedinputs

import (
	"github.com/andrewhowdencom/bmctl/cmd/errors"
	"github.com/spf13/cobra"
)

// SupportedInputsCmd represents the audio supported-inputs command
var SupportedInputsCmd = &cobra.Command{
	Use:   "supported-inputs",
	Short: "Commands associated with managing the audio supported inputs",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.ErrNoSubcommand
	},
}

func init() {
	SupportedInputsCmd.AddCommand(GetSupportedInputsCmd)
}
