package cmd

import (
	"errors"
	"fmt"
	"timetracker_cli/internal"

	"github.com/spf13/cobra"
)

func NewStartCmd(tracker *internal.Tracker) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start tracking time",
		Long:  `Starts a new tracking session only if session is not in progress`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if tracker.Current.Start > 0 {
				fmt.Fprint(cmd.OutOrStdout(), "Session has already started")
				return nil
			}

			if len(args) > 1 {
				return errors.New("start only accepts 0 or 1 arguments")
			}

			if len(args) == 1 {
				tracker.SetDescription(args[0])
			}

			tracker.Start()

			return nil
		},
	}
}
