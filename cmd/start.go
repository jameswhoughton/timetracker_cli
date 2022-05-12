package cmd

import (
	"fmt"
	"timetracker_cli/internal"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
func NewStartCmd(tracker *internal.Tracker) *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start tracking time",
		Long:  `Starts a new tracking session only if session is not in progress`,
		Run: func(cmd *cobra.Command, args []string) {
			if tracker.Current.Start > 0 {
				fmt.Fprint(cmd.OutOrStdout(), "Session has already started")
				return
			}

			tracker.Start()
		},
	}
}

// func init() {
// 	rootCmd.AddCommand(startCmd)

// 	// Here you will define your flags and configuration settings.

// 	// Cobra supports Persistent Flags which will work for this command
// 	// and all subcommands, e.g.:
// 	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	// Cobra supports local flags which will only run when this command
// 	// is called directly, e.g.:
// 	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
