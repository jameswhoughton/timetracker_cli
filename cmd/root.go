/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"timetracker_cli/internal"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
func NewRootCmd(tracker *internal.Tracker) *cobra.Command {
	return &cobra.Command{
		Use:   "timetracker_cli",
		Short: "Track your time",
		Long:  `Track your time`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			if len(tracker.Sessions) == 0 {
				fmt.Fprintf(cmd.OutOrStdout(), "No Sessions")
			}

			for i, total := range tracker.Totals {
				fmt.Fprintf(cmd.OutOrStdout(), "%s - %s", total.Description, internal.FormatTotal(total.Total))
				if i < len(tracker.Totals)-1 {
					fmt.Fprint(cmd.OutOrStdout(), "\n")
				}
			}
		},
	}
}

// // Execute adds all child commands to the root command and sets flags appropriately.
// // This is called by main.main(). It only needs to happen once to the rootCmd.
// func Execute() {
// 	err := rootCmd.Execute()
// 	if err != nil {
// 		os.Exit(1)
// 	}
// }

// func init() {
// 	// Here you will define your flags and configuration settings.
// 	// Cobra supports persistent flags, which, if defined here,
// 	// will be global for your application.

// 	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.timetracker_cli.yaml)")

// 	// Cobra also supports local flags, which will only run
// 	// when this action is called directly.
// 	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
