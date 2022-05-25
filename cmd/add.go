package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		addInt(args)
	},
}

func addInt(args []string) {
	var sum int
	for _, value := range args {
		temp, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println(err)
		}
		sum = sum + temp
	}
	fmt.Printf("Sum of numbers: %s is %d\n", args, sum)
}

func init() {
	rootCmd.AddCommand(addCmd)
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
