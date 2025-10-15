/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/jamesfulreader/gpcli/internal/cli"
	"github.com/jamesfulreader/gpcli/internal/tsdb"
	"github.com/spf13/cobra"
)

var (
	app = &cli.App{
		Options: cli.AppOptions{},
		Out:     nil,
		Client:  nil,
	}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gpcli",
	Short: "CLI for F1/F3 data via TheSportsDB",
	Long:  `Query Formula 1 and Formula 3 data from TheSportsDB API.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		app.Options.APIKey = "123"
		app.Out = cmd.OutOrStdout()
		app.Client = tsdb.NewClient(app.Options.APIKey)
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gpcli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVar(
		&app.Options.APIKey,
		"api-key",
		"",
		"Default free SportsDB API key 123",
	)
	rootCmd.PersistentFlags().StringVar(
		&app.Options.League,
		"league",
		"f1",
		"League: f1, f3, or numeric ID",
	)
	rootCmd.PersistentFlags().BoolVar(
		&app.Options.JSON,
		"json",
		false,
		"Output JSON instead of text",
	)
}
