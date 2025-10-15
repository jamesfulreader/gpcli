/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"time"

	icl "github.com/jamesfulreader/gpcli/internal/cli"
	"github.com/spf13/cobra"
)

// leaguesCmd represents the leagues command
var leaguesCmd = &cobra.Command{
	Use:   "leagues",
	Short: "Show league info (use --league f1|f3 or numeric ID)",

	RunE: func(cmd *cobra.Command, args []string) error {
		leagueID, err := app.ResolveLeagueID()
		if err != nil {
			return err
		}
		ctx, cancel := context.WithTimeout(cmd.Context(), 10*time.Second)
		defer cancel()
		return icl.ShowLeague(ctx, app, leagueID)
	},
}

func init() {
	rootCmd.AddCommand(leaguesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// leaguesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// leaguesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
