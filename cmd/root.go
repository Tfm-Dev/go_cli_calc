package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "calc",
	Short: "Hello Calculator",
	Long:  `This a calculator cli using cobra`,
	Run:   func(cmd *cobra.Command, args []string) { fmt.Println("This a calculator with cobra cli!") },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
