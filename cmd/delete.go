/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"strconv"
	"timetracker_cli/internal"

	"github.com/spf13/cobra"
)

func NewDeleteCmd(tracker *internal.Tracker) *cobra.Command {
	return &cobra.Command{
		Use:   "delete",
		Short: "Delete a session",
		Long:  `Delete a session by ID`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("delete expects 1 argument: the ID of the session to delete")
			}

			id, _ := strconv.Atoi(args[0])
			index := id - 1

			err := tracker.DeleteByIndex(index)

			if err != nil {
				return err
			}

			return nil
		},
	}
}
