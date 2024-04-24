package cmd

import (
	"github.com/andrewhowdencom/bmctl/cmd/errors"
	"github.com/andrewhowdencom/bmctl/cmd/lens"
	"github.com/spf13/cobra"
)

// lensCmd represents the lens command
var lensCmd = &cobra.Command{
	Use:   "lens",
	Short: "Commands associated with managing the lens",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.ErrNoSubcommand
	},
}

func init() {
	lensCmd.AddCommand(lens.AutoFocusCmd)
}
