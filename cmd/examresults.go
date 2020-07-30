package cmd

import (
	"fmt"
	"os"
	"strings"

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

		exams, err := acc.Examresults(semester)
		if err != nil {
			fmt.Printf("Unable to export: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("%d exam results:\n", len(exams))
		for _, exam := range exams {
			fmt.Println(strings.Join(exam, ", "))
		}
	},
}

func init() {
	rootCmd.AddCommand(examresultsCmd)

	examresultsCmd.Flags().StringVarP(&semester, "semester", "s", "", "Semester ID (eg. '099999904632582' for SoSe20, '999' for all; defaults to current semester)")
}
