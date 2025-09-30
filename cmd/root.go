package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/andrewhowdencom/bmctl/cmd/errors"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bmctl",
	Short: "Black Magic Control Utility",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if !viper.IsSet("api.server") {
			return fmt.Errorf("api.server must be set in config, or via the --api.server flag")
		}
		return nil
	},
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
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP("api.server", "a", "", "The HTTP(S) address of the camera API Server")
	viper.BindPFlag("api.server", rootCmd.PersistentFlags().Lookup("api.server"))

	rootCmd.AddCommand(lensCmd)
	rootCmd.AddCommand(videoCmd)
	rootCmd.AddCommand(audioCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			logger.Error("could not find home directory", "error", err)
			os.Exit(1)
		}

		// Search config in home directory with name ".bmctl" (without extension).
		viper.AddConfigPath(fmt.Sprintf("%s/.config/bmctl", home))
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; prompt user to create one.
			logger.Info("no config file found, consider creating one at ~/.config/bmctl/config.yaml")
		} else {
			// A non-trivial error occurred, so we should probably exit.
			logger.Error("could not read config file", "error", err)
			os.Exit(1)
		}
	}
}
