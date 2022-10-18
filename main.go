package main

import (
	"os"
	"timetracker_cli/cmd"
	"timetracker_cli/internal"
)

func setup(tracker *internal.Tracker, config config) {
	rootCmd := cmd.NewRootCmd()

	listCmd := cmd.NewListCmd(tracker, cmd.ListConfig{RoundBy: config.RoundBy})
	totalsCmd := cmd.NewTotalsCmd(tracker, cmd.TotalsConfig{RoundBy: config.RoundBy})
	startCmd := cmd.NewStartCmd(tracker)
	endCmd := cmd.NewEndCmd(tracker)
	deletetCmd := cmd.NewDeleteCmd(tracker)
	addCmd := cmd.NewAddCmd(tracker)

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(totalsCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(endCmd)
	rootCmd.AddCommand(deletetCmd)
	rootCmd.AddCommand(addCmd)

	rootCmd.Execute()
}

var sessionFile string
var roundBy int

type config struct {
	SessionFile string
	RoundBy     int
}

func loadConfig() config {
	if sessionFile == "" {
		userDir, _ := os.UserConfigDir()

		if _, err := os.Stat(userDir + "/tt"); err != nil {
			os.Mkdir(userDir+"/tt", os.ModePerm)
		}

		sessionFile = userDir + "/tt/session.json"
	}

	if roundBy == 0 {
		roundBy = 900
	}

	return config{
		SessionFile: sessionFile,
		RoundBy:     roundBy,
	}
}

func main() {
	config := loadConfig()

	tracker := internal.NewTracker(internal.TrackerConfig{File: config.SessionFile})
	defer tracker.Save()

	setup(tracker, config)
}
