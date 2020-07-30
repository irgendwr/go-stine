package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// examsCmd represents the exams command
var examsCmd = &cobra.Command{
	Use:   "exams",
	Short: "List of exams",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		acc, err := login()
		if err != nil {
			fmt.Printf("Unable to log-in: %s\n", err)
			os.Exit(1)
		}

		exams, err := acc.Exams(semester)
		if err != nil {
			fmt.Printf("Unable to export: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("%d exams:\n", len(exams))
		for _, exam := range exams {
			fmt.Println(strings.Join(exam, ", "))
		}
	},
}

func init() {
	rootCmd.AddCommand(examsCmd)

	examsCmd.Flags().StringVarP(&semester, "semester", "s", "", "Semester ID (eg. '099999904632582' for SoSe20, '999' for all; defaults to current semester)")
}
