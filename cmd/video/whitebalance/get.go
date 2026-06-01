package whitebalance

import (
	"fmt"
	"os"

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
	opts := []client.Option{}
	if viper.GetBool("curl") {
		opts = append(opts, client.WithCurlOutput(os.Stderr))
	}

	client, err := client.New(viper.GetString("api.server"), opts...)
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
