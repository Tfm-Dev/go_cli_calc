package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "go_cli_calc",
	Long: ``,
	Run:  func(cmd *cobra.Command, args []string) { fmt.Println("This a calc with cobra cli!") },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
