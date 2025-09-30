package audio

import (
	"github.com/andrewhowdencom/bmctl/cmd/audio/input"
	"github.com/andrewhowdencom/bmctl/cmd/errors"
	"github.com/spf13/cobra"
)

// InputCmd represents the audio input command
var InputCmd = &cobra.Command{
	Use:   "input",
	Short: "Commands associated with managing the audio input",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.ErrNoSubcommand
	},
}

func init() {
	InputCmd.AddCommand(input.GetInputCmd)
	InputCmd.AddCommand(input.SetInputCmd)
}