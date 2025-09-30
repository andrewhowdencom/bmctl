package whitebalance

import (
	"fmt"

	"github.com/andrewhowdencom/bmctl/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetWhiteBalanceCmd represents the lens command
var GetWhiteBalanceCmd = &cobra.Command{
	Use:   "get",
	Short: "Get white balance for the video stream",
	RunE:  DoGetWhiteBalance,
}

func DoGetWhiteBalance(cmd *cobra.Command, args []string) error {
	client, err := client.New(viper.GetString("api.server"))
	if err != nil {
		return err
	}

	i, err := client.VideoGetWhiteBalance()
	if err != nil {
		return err
	}

	fmt.Println(i)

	return nil
}
