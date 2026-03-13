package available

import (
	"github.com/andrewhowdencom/bmctl/cmd/errors"
	"github.com/spf13/cobra"
)

// AvailableCmd represents the audio available command
var AvailableCmd = &cobra.Command{
	Use:   "available",
	Short: "Commands associated with managing the audio available status",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.ErrNoSubcommand
	},
}

func init() {
	AvailableCmd.AddCommand(GetAvailableCmd)
}
