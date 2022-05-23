package cmd

import (
	"fmt"
	"timetracker_cli/internal"

	"github.com/spf13/cobra"
)

func NewRootCmd(tracker *internal.Tracker) *cobra.Command {
	return &cobra.Command{
		Use:   "timetracker_cli",
		Short: "Track your time",
		Long:  `Track your time`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(tracker.Sessions) == 0 {
				fmt.Fprintf(cmd.OutOrStdout(), "No sessions")
			}

			for i, total := range tracker.Totals {
				fmt.Fprintf(cmd.OutOrStdout(), "%s - %s", total.Description, internal.FormatTotal(total.Total))
				if i < len(tracker.Totals)-1 {
					fmt.Fprint(cmd.OutOrStdout(), "\n")
				}
			}
		},
	}
}
