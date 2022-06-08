package main

import (
	"timetracker_cli/cmd"
	"timetracker_cli/internal"
)

func setup(tracker *internal.Tracker) {
	rootCmd := cmd.NewRootCmd(tracker)

	listCmd := cmd.NewListCmd(tracker)
	startCmd := cmd.NewStartCmd(tracker)
	endCmd := cmd.NewEndCmd(tracker)
	deletetCmd := cmd.NewDeleteCmd(tracker)

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(endCmd)
	rootCmd.AddCommand(deletetCmd)

	rootCmd.Execute()
}

func main() {
	config := internal.TrackerConfig{
		File: "session.json",
	}

	tracker := internal.NewTracker(config)
	defer tracker.Save()

	setup(tracker)
}
