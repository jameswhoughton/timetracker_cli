package cmd

import (
	"fmt"
	"timetracker_cli/internal"

	"github.com/rodaine/table"
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
				return
			}

			sessionTable := table.New("Description", "Total")

			sessionTable.WithWriter(cmd.OutOrStdout())

			for _, total := range tracker.Totals {
				sessionTable.AddRow(total.Description, internal.FormatTotal(total.Total))
			}

			sessionTable.Print()
		},
	}
}
