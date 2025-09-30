package lens

import (
	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// AutoFocusCmd represents the lens command
var AutoFocusCmd = &cobra.Command{
	Use:   "autofocus",
	Short: "Automatically focus the lens",
	RunE:  DoAutoFocusCmd,
}

func DoAutoFocusCmd(cmd *cobra.Command, args []string) error {
	client, err := client.New(viper.GetString("api.server"))
	if err != nil {
		return err
	}

	if err := client.LensAutoFocus(); err != nil {
		return err
	}

	return nil
}
