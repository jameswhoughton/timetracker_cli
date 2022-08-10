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

			sessionTable := table.New("ID", "Description", "Start", "End", "Length")

			sessionTable.WithWriter(cmd.OutOrStdout())

			for id, session := range tracker.Sessions {
				start := internal.FormatTime(session.Start, "02/01/06 15:04:05")
				end := internal.FormatTime(session.End, "02/01/06 15:04:05")
				length := internal.FormatTotal(int(session.End - session.Start))
				sessionTable.AddRow(id+1, session.Description, start, end, length)
			}

			sessionTable.Print()
		},
	}
}
