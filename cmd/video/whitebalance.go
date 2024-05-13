package video

import (
	"github.com/andrewhowdencom/bmctl/cmd/errors"
	"github.com/andrewhowdencom/bmctl/cmd/video/whitebalance"
	"github.com/spf13/cobra"
)

// isoCmd represents the lens command
var WhiteBalanceCMD = &cobra.Command{
	Use:   "whitebalance",
	Short: "Commands associated with managing the whitebalance",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.ErrNoSubcommand
	},
}

func init() {
	WhiteBalanceCMD.AddCommand(whitebalance.GetWhiteBalanceCmd)
	WhiteBalanceCMD.AddCommand(whitebalance.SetWhiteBalanceCmd)
}
