package cmd

import (
	"context"
	"time"

	icl "github.com/jamesfulreader/gpcli/internal/cli"
	"github.com/spf13/cobra"
)

var eventID string

var eventCmd = &cobra.Command{
	Use:   "event [id]",
	Short: "Show details for a specific event by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), 10*time.Second)
		defer cancel()
		return icl.ShowEvent(ctx, app, args[0])
	},
}

func init() {
	rootCmd.AddCommand(eventCmd)
	// flag form (optional, since we allow positional)
	eventCmd.Flags().StringVar(&eventID, "id", "", "Event ID")
}
