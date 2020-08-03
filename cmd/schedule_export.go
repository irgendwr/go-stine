package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// scheduleExportCmd represents the export command
var scheduleExportCmd = &cobra.Command{
	Use:   "export <date>",
	Short: "Exports the schedule of a given month/week",
	Long: `Exports the schedule of a given month or week.

<date> can be either a month (e.g. Y2020M06) or a week (e.g. Y2020W25).
By default the schedule will be exported as an UTF-16 LE encoded .ics file named <date>.ics in your CWD.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("expected a date as argument, received %d arguments instead", len(args))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			acc, err := login()
			if err != nil {
				fmt.Printf("Unable to log-in: %s\n", err)
				os.Exit(1)
			}

			date := args[0]
			path := output
			if path == "" {
				path = date + ".ics"
			}

			url, err := acc.SchedulerExport(date)
			if err != nil {
				fmt.Printf("Unable to export: %s\n", err)
				os.Exit(1)
			}

			if err := DownloadFile(acc, path, url); err != nil {
				fmt.Printf("Unable to download: %s\n", err)
				os.Exit(1)
			}
		} else {
			cmd.Usage()
		}
	},
}

func init() {
	scheduleCmd.AddCommand(scheduleExportCmd)

	scheduleExportCmd.Flags().StringVarP(&output, "output", "o", "", "output file (default: <date>.ics)")
}
