package cmd

import (
	"github.com/andrewhowdencom/bmctl/cmd/errors"
	"github.com/andrewhowdencom/bmctl/cmd/video"
	"github.com/spf13/cobra"
)

// videoCmd represents the lens command
var videoCmd = &cobra.Command{
	Use:   "video",
	Short: "Commands associated with managing the video",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.ErrNoSubcommand
	},
}

func init() {
	videoCmd.AddCommand(video.ISOCmd)
}
