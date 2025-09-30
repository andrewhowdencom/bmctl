package whitebalance

import (
	"fmt"
	"strconv"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/andrewhowdencom/bmctl/cmd/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// SetWhiteBalanceCmd sets the white balance for command
var SetWhiteBalanceCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the white balance for the video stream",
	Args:  cobra.ExactArgs(1),
	RunE:  DoSetWhiteBalance,
}

func DoSetWhiteBalance(cmd *cobra.Command, args []string) error {
	n, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w: %s", errors.ErrNotANumber, err)
	}

	client, err := client.New(viper.GetString("api.server"))
	if err != nil {
		return err
	}

	if err := client.VideoSetWhiteBalance(n); err != nil {
		return err
	}

	return nil
}
