package lowcutfilter

import (
	"github.com/andrewhowdencom/bmctl/cmd/errors"
	"github.com/spf13/cobra"
)

// LowCutFilterCmd represents the audio low-cut-filter command
var LowCutFilterCmd = &cobra.Command{
	Use:   "low-cut-filter",
	Short: "Commands associated with managing the audio low cut filter",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.ErrNoSubcommand
	},
}

func init() {
	LowCutFilterCmd.AddCommand(GetLowCutFilterCmd)
	LowCutFilterCmd.AddCommand(SetLowCutFilterCmd)
}
