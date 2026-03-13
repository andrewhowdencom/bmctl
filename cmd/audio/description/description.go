package description

import (
	"github.com/andrewhowdencom/bmctl/cmd/errors"
	"github.com/spf13/cobra"
)

// DescriptionCmd represents the audio description command
var DescriptionCmd = &cobra.Command{
	Use:   "description",
	Short: "Commands associated with managing the audio description",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.ErrNoSubcommand
	},
}

func init() {
	DescriptionCmd.AddCommand(GetDescriptionCmd)
}
