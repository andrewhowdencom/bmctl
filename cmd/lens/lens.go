package lens

import (
	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
)

// AutoFocusCmd represents the lens command
var AutoFocusCmd = &cobra.Command{
	Use:   "autofocus",
	Short: "Automatically focus the lens",
	RunE:  DoAutoFocusCmd,
}

func DoAutoFocusCmd(cmd *cobra.Command, args []string) error {
	client, err := client.New(cmd.Flags().Lookup("api.server").Value.String())
	if err != nil {
		return err
	}

	if err := client.LensAutoFocus(); err != nil {
		return err
	}

	return nil
}
