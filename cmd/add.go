package cmd

import (
	"strconv"
	"strings"
	"time"
	"timetracker_cli/internal"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
func NewAddCmd(tracker *internal.Tracker) *cobra.Command {
	var start string
	var end string
	var description string

	cmd := cobra.Command{
		Use:   "add",
		Short: "Manually add a session",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			if start == "" {
				start = internal.TimeInput("start")
			}

			if end == "" {
				end = internal.TimeInput("end")
			}

			if description == "" {
				description = internal.DescriptionInput()
			}

			session := internal.Session{
				Start:       timeToday(start).Unix(),
				End:         timeToday(end).Unix(),
				Description: description,
			}

			tracker.Add(session)
		},
	}

	cmd.Flags().StringVar(&start, "start", "", "start")
	cmd.Flags().StringVar(&end, "end", "", "end")
	cmd.Flags().StringVar(&description, "description", "", "description")

	return &cmd
}

func timeToday(timeString string) time.Time {
	today := time.Now()

	timeComponents := strings.Split(timeString, ":")
	hours, _ := strconv.Atoi(timeComponents[0])
	minutes, _ := strconv.Atoi(timeComponents[1])

	return time.Date(today.Year(), today.Month(), today.Day(), hours, minutes, 0, 0, today.Location())
}
