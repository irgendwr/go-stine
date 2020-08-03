package cmd

import (
	"github.com/spf13/cobra"
)

// schedulerCmd represents the schedule command
var scheduleCmd = &cobra.Command{
	Use:     "schedule",
	Aliases: []string{"scheduler"},
	Short:   "group of schedule commands",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(scheduleCmd)
}
