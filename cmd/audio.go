package cmd

import (
	"github.com/andrewhowdencom/bmctl/cmd/audio"
	"github.com/andrewhowdencom/bmctl/cmd/audio/available"
	"github.com/andrewhowdencom/bmctl/cmd/audio/description"
	"github.com/andrewhowdencom/bmctl/cmd/audio/level"
	"github.com/andrewhowdencom/bmctl/cmd/audio/lowcutfilter"
	"github.com/andrewhowdencom/bmctl/cmd/audio/padding"
	"github.com/andrewhowdencom/bmctl/cmd/audio/phantompower"
	"github.com/andrewhowdencom/bmctl/cmd/audio/supportedinputs"
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
	audioCmd.AddCommand(available.AvailableCmd)
	audioCmd.AddCommand(description.DescriptionCmd)
	audioCmd.AddCommand(level.LevelCmd)
	audioCmd.AddCommand(lowcutfilter.LowCutFilterCmd)
	audioCmd.AddCommand(padding.PaddingCmd)
	audioCmd.AddCommand(phantompower.PhantomPowerCmd)
	audioCmd.AddCommand(supportedinputs.SupportedInputsCmd)
}