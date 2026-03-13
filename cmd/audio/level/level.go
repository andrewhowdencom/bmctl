package level

import (
	"github.com/andrewhowdencom/bmctl/cmd/errors"
	"github.com/spf13/cobra"
)

// LevelCmd represents the audio level command
var LevelCmd = &cobra.Command{
	Use:   "level",
	Short: "Commands associated with managing the audio level",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.ErrNoSubcommand
	},
}

func init() {
	LevelCmd.AddCommand(GetLevelCmd)
	LevelCmd.AddCommand(SetLevelCmd)
}
