package cmd

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/irgendwr/table"
	"github.com/spf13/cobra"
)

// examsCmd represents the exams command
var examresultsCmd = &cobra.Command{
	Use:   "examresults",
	Short: "List of exam results",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		acc, err := login()
		if err != nil {
			fmt.Printf("Unable to log-in: %s\n", err)
			os.Exit(1)
		}

		if all, err := cmd.Flags().GetBool("all"); err == nil && all {
			semester = "999"
		}

		exams, err := acc.ExamResults(semester)
		if err != nil {
			fmt.Printf("Unable to export: %s\n", err)
			os.Exit(1)
		}

		if fcsv, err := cmd.Flags().GetBool("csv"); err == nil && fcsv {
			csvWriter := csv.NewWriter(os.Stdout)
			if err = csvWriter.WriteAll(exams); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			return
		}

		tbl := table.New("ID", "Name", "Date", "Grade", "")
		tbl.WithHeaderFormatter(color.New(color.Underline).SprintfFunc())
		tbl.SetRows(exams)
		tbl.Print()
	},
}

func init() {
	rootCmd.AddCommand(examresultsCmd)

	examresultsCmd.Flags().StringVarP(&semester, "semester", "s", "", "Semester ID (eg. '099999904632582' for SoSe20; defaults to current semester)")
	examresultsCmd.Flags().BoolP("all", "a", false, "Selects all semesters")
	examresultsCmd.Flags().Bool("csv", false, "Outputs CSV instead of table")
}
