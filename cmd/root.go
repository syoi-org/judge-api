package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/syoi-org/judge-api/config"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "judy",
	Short: "Command line utility for running Online Judge API",
	Long: `This is a command line utility for running Online Judge API.

This API is a RESTful API that provides a set of endpoints for submitting and
retrieving problems, submissions, and judgements to popular online judge
systems.
For more information, see https://github.com/syoi-org/judge-api`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		config.InitConfig(configFile)
	})
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "path to config file")
}
