package cmd

import (
	"os"

	"github.com/andrewhowdencom/bmctl/cmd/errors"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bmctl",
	Short: "Black Magic Control Utility",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.ErrNoSubcommand
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("api.server", "a", "", "The HTTP(S) address of the camera API Server")
	rootCmd.MarkPersistentFlagRequired("api.server")

	rootCmd.AddCommand(lensCmd)
	rootCmd.AddCommand(videoCmd)
}
