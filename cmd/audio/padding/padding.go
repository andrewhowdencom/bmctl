package padding

import (
	"github.com/andrewhowdencom/bmctl/cmd/errors"
	"github.com/spf13/cobra"
)

// PaddingCmd represents the audio padding command
var PaddingCmd = &cobra.Command{
	Use:   "padding",
	Short: "Commands associated with managing the audio padding",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.ErrNoSubcommand
	},
}

func init() {
	PaddingCmd.AddCommand(GetPaddingCmd)
	PaddingCmd.AddCommand(SetPaddingCmd)
}
