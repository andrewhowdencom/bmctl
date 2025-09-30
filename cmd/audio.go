package cmd

import (
	"github.com/andrewhowdencom/bmctl/cmd/audio"
	"github.com/andrewhowdencom/bmctl/cmd/errors"
	"github.com/spf13/cobra"
)

// audioCmd represents the audio command
var audioCmd = &cobra.Command{
	Use:   "audio",
	Short: "Commands associated with managing the audio",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.ErrNoSubcommand
	},
}

func init() {
	audioCmd.AddCommand(audio.InputCmd)
}