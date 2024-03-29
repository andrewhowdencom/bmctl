package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

var ErrNoSubcommand = errors.New("no subcommand supplied")

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bmctl",
	Short: "Black Magic Control Utility",
	RunE: func(cmd *cobra.Command, args []string) error {
		return ErrNoSubcommand
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
