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

var scheduleSeason string

// scheduleCmd represents the schedule command
var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Show the season schedule for the selected league",
	RunE: func(cmd *cobra.Command, args []string) error {
		leagueID, err := app.ResolveLeagueID()
		if err != nil {
			return err
		}
		ctx, cancel := context.WithTimeout(cmd.Context(), 10*time.Second)
		defer cancel()
		return icl.ShowSchedule(ctx, app, leagueID, scheduleSeason)
	},
}

func init() {
	rootCmd.AddCommand(scheduleCmd)

	scheduleCmd.Flags().StringVar(
		&scheduleSeason,
		"season",
		"2025",
		"Season year (e.g., 2025)",
	)
}
