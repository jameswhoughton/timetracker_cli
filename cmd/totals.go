package cmd

import (
	"fmt"
	"timetracker_cli/internal"

	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

type TotalsConfig struct {
	RoundBy int
}

// totalsCmd represents the totals command
func NewTotalsCmd(tracker *internal.Tracker, config TotalsConfig) *cobra.Command {
	return &cobra.Command{
		Use:   "totals",
		Short: "List the grouped totals",
		Long:  "Each entry is grouped by the name, totals displays the total time for each group.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(tracker.Sessions) == 0 {
				fmt.Fprintf(cmd.OutOrStdout(), "No sessions")
				return
			}

			headers := []interface{}{"Description", "Total"}

			if config.RoundBy > 0 {
				headers = append(headers, "Rounded")
			}

			sessionTable := table.New(headers...)
			sessionTable.WithWriter(cmd.OutOrStdout())

			for _, session := range tracker.Totals {
				row := []interface{}{
					session.Description,
					internal.FormatTotal(session.Total),
				}

				if config.RoundBy > 0 {
					row = append(row, internal.FormatTotal(internal.Round(session.Total, config.RoundBy)))
				}

				sessionTable.AddRow(row...)
			}

			sessionTable.Print()
		},
	}
}
