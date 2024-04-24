package video

import (
	"github.com/andrewhowdencom/bmctl/cmd/errors"
	"github.com/andrewhowdencom/bmctl/cmd/video/iso"
	"github.com/spf13/cobra"
)

// isoCmd represents the lens command
var ISOCmd = &cobra.Command{
	Use:   "iso",
	Short: "Commands associated with managing the iso",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.ErrNoSubcommand
	},
}

func init() {
	ISOCmd.AddCommand(iso.GetISOCmd)
	ISOCmd.AddCommand(iso.SetISOCmd)
}
