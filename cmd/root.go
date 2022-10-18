package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "timetracker_cli",
		Short: "Track your time",
		Long:  `Track your time`,
	}
}
