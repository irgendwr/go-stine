package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/irgendwr/go-stine/api"
	"github.com/irgendwr/table"
	"github.com/spf13/cobra"
)

// schedulerCmd represents the schedule command
var scheduleCmd = &cobra.Command{
	Use:     "schedule [date]",
	Aliases: []string{"scheduler"},
	Short:   "group of schedule commands",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		acc, err := login()
		if err != nil {
			fmt.Printf("Unable to log-in: %s\n", err)
			os.Exit(1)
		}

		date := ""
		if len(args) > 0 {
			date = args[0]
		}

		skip := api.SkipNone
		if prev, err := cmd.Flags().GetBool("prev"); err == nil && prev {
			skip = api.SkipPrev
		}
		if next, err := cmd.Flags().GetBool("next"); err == nil && next {
			skip = api.SkipNext
		}

		view := api.ScheduleWeek
		if day, err := cmd.Flags().GetBool("day"); err == nil && day {
			view = api.ScheduleDay
		}

		schedules, err := acc.Scheduler(date, skip, view)
		if err != nil {
			fmt.Printf("Unable to export: %s\n", err)
			os.Exit(1)
		}

		for i, schedule := range schedules {
			if i != 0 {
				fmt.Println()
			}

			color.New(color.Bold).Println(schedule.Date)
			tbl := table.New("ID", "Name", "Teachers", "Time", "Room")
			tbl.WithHeaderFormatter(color.New(color.Underline).SprintfFunc())
			tbl.SetRows(schedule.Entries)
			tbl.Print()
		}
	},
}

func init() {
	rootCmd.AddCommand(scheduleCmd)

	scheduleCmd.Flags().Bool("prev", false, "Previous (does nothing if --next is set)")
	scheduleCmd.Flags().Bool("next", false, "Next")
	scheduleCmd.Flags().Bool("week", false, "Week (does nothing if --day is set)")
	scheduleCmd.Flags().Bool("day", false, "Day")
}
