package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "Subtration calculation",
	Long: `Subtration calculation with numbers integer or floating.
Need two or more arguments to be added.`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Calculator subtration")
	},
}

func init() {
	rootCmd.AddCommand(subCmd)
}
