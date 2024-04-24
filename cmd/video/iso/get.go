package iso

import (
	"fmt"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
)

// GetISOCmd represents the lens command
var GetISOCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the ISO for the video stream",
	RunE:  DoGetISO,
}

func DoGetISO(cmd *cobra.Command, args []string) error {
	client, err := client.New(cmd.Flags().Lookup("api.server").Value.String())
	if err != nil {
		return err
	}

	i, err := client.VideoGetISO()
	if err != nil {
		return err
	}

	fmt.Println(i)

	return nil
}
