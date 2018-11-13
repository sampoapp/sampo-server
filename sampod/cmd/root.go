/* This is free and unencumbered software released into the public domain. */

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string
var debug bool
var verbose bool

// RootCmd describes the `sampod` command
var RootCmd = &cobra.Command{
	Use:   "sampod",
	Short: "Sampo daemon",
	Long: `Sampo is a personal information manager (PIM) app.
This is the server daemon for Sampo.`,
	Version: "0.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO") // TODO
	},
}

// Execute implements the `sampod` command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVarP(&configFile, "config", "C", "", "Set config file (default: $HOME/.sampo/config.yaml)")
	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debugging")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Be verbose")
	RootCmd.SetVersionTemplate(`Sampo daemon {{printf "%s" .Version}}
`)
}

// initConfig reads the configuration file and environment variables.
func initConfig() {
	if configFile != "" {
		// Use the specified config file:
		viper.SetConfigFile(configFile)
	} else {
		// Search for config file in the current directory and under the home directory:
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.sampo")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in:
	if err := viper.ReadInConfig(); err == nil {
		if debug {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}
}
