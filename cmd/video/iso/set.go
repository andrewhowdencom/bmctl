package iso

import (
	"fmt"
	"strconv"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/andrewhowdencom/bmctl/cmd/errors"
	"github.com/spf13/cobra"
)

// SetISOCmd represents the lens command
var SetISOCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the ISO for the video stream",
	Args:  cobra.ExactArgs(1),
	RunE:  DoSetISO,
}

func DoSetISO(cmd *cobra.Command, args []string) error {
	n, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w: %s", errors.ErrNotANumber, err)
	}

	client, err := client.New(cmd.Flags().Lookup("api.server").Value.String())
	if err != nil {
		return err
	}

	if err := client.VideoSetISO(n); err != nil {
		return err
	}

	return nil
}
