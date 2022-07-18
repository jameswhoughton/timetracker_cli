package main

import (
	"encoding/json"
	"os"
	"timetracker_cli/cmd"
	"timetracker_cli/internal"
)

func setup(tracker *internal.Tracker) {
	rootCmd := cmd.NewRootCmd(tracker)

	listCmd := cmd.NewListCmd(tracker)
	startCmd := cmd.NewStartCmd(tracker)
	endCmd := cmd.NewEndCmd(tracker)
	deletetCmd := cmd.NewDeleteCmd(tracker)
	addCmd := cmd.NewAddCmd(tracker)

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(endCmd)
	rootCmd.AddCommand(deletetCmd)
	rootCmd.AddCommand(addCmd)

	rootCmd.Execute()
}

var config_path string

func main() {
	configFile, err := os.ReadFile(config_path)

	if err != nil {
		panic(err)
	}

	config := internal.TrackerConfig{}

	json.Unmarshal(configFile, &config)

	tracker := internal.NewTracker(config)
	defer tracker.Save()

	setup(tracker)
}
