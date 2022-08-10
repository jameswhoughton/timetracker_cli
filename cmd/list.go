package cmd

import (
	"fmt"
	"timetracker_cli/internal"

	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

type ListConfig struct {
	RoundBy int
}

func NewListCmd(tracker *internal.Tracker, config ListConfig) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all sessions",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if len(tracker.Sessions) == 0 {
				fmt.Fprintf(cmd.OutOrStdout(), "No sessions")
				return
			}

			headers := []interface{}{"ID", "Description", "Start", "End", "Length"}

			if config.RoundBy > 0 {
				headers = append(headers, "Rounded")
			}

			sessionTable := table.New(headers...)
			sessionTable.WithWriter(cmd.OutOrStdout())

			for id, session := range tracker.Sessions {
				row := []interface{}{
					id + 1,
					session.Description,
					internal.FormatTime(session.Start, "02/01/06 15:04:05"),
					internal.FormatTime(session.End, "02/01/06 15:04:05"),
					internal.FormatTotal(int(session.End - session.Start)),
				}

				if config.RoundBy > 0 {
					row = append(row, internal.FormatTotal(internal.Round(int(session.End-session.Start), config.RoundBy)))
				}
				sessionTable.AddRow(row...)
			}

			sessionTable.Print()
		},
	}
}
