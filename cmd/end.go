package cmd

import (
	"errors"
	"timetracker_cli/internal"

	"github.com/spf13/cobra"
)

func NewEndCmd(tracker *internal.Tracker) *cobra.Command {
	return &cobra.Command{
		Use:   "end",
		Short: "End the current session",
		Long:  `End the current session with an optional description`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if tracker.Current.Start == 0 {
				return errors.New("cannot end a session that hasn't started")
			}

			if len(args) > 1 {
				return errors.New("end only accepts 0 or 1 arguments")
			}

			if len(args) == 1 {
				tracker.SetDescription(args[0])
			}

			if tracker.Current.Description == "" {
				return errors.New("cannot end a session without adding a description")
			}

			return tracker.End()
		},
	}
}
