package main

import (
	"encoding/json"
	"os"
	"timetracker_cli/cmd"
	"timetracker_cli/internal"
)

func setup(tracker *internal.Tracker, config config) {
	rootCmd := cmd.NewRootCmd(tracker)

	listCmd := cmd.NewListCmd(tracker, cmd.ListConfig{RoundBy: config.RoundBy})
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

type config struct {
	SessionFile string `json:session_file`
	RoundBy     int    `json:round_by`
}

func main() {
	configFile, err := os.ReadFile(config_path)

	if err != nil {
		panic(err)
	}

	config := config{}

	json.Unmarshal(configFile, &config)

	tracker := internal.NewTracker(internal.TrackerConfig{File: config.SessionFile})
	defer tracker.Save()

	setup(tracker, config)
}
