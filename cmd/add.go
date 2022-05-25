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
		float, _ := cmd.Flags().GetBool("float")
		if float {
			addFloat(args)
		} else {
			addInt(args)
		}
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

func addFloat(args []string) {
	var sum float64
	for _, value := range args {
		temp, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Println(err)
		}
		sum = sum + temp
	}
	fmt.Printf("Sum of numbers: %s is %.2f\n", args, sum)
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("float", "f", false, "add floating numbers")
}
