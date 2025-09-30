package iso

import (
	"fmt"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetISOCmd represents the lens command
var GetISOCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the ISO for the video stream",
	RunE:  DoGetISO,
}

func DoGetISO(cmd *cobra.Command, args []string) error {
	client, err := client.New(viper.GetString("api.server"))
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
