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

// teamsCmd represents the teams command
var teamsCmd = &cobra.Command{
	Use:   "teams",
	Short: "List teams/constructors for the selected league",

	RunE: func(cmd *cobra.Command, args []string) error {
		leagueID, err := app.ResolveLeagueID()
		if err != nil {
			return err
		}
		ctx, cancel := context.WithTimeout(cmd.Context(), 10*time.Second)
		defer cancel()

		return icl.ListTeams(ctx, app, leagueID)
	},
}

func init() {
	rootCmd.AddCommand(teamsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// teamsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// teamsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
