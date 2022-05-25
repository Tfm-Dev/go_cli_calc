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
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		float, _ := cmd.Flags().GetBool("float")
		if float {
			addFloat(args)
		} else {
			odd, _ := cmd.Flags().GetBool("odd")
			even, _ := cmd.Flags().GetBool("even")
			addInt(args, odd, even)
		}
	},
}

func addInt(args []string, odd bool, even bool) {
	var sum int
	for _, value := range args {
		temp, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println(err)
		}
		if odd && temp%2 == 0 || even && temp%2 != 0 || !odd && !even {
			sum = sum + temp
		}
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
	addCmd.Flags().BoolP("odd", "o", false, "calculate with odd numbers")
	addCmd.Flags().BoolP("even", "e", false, "calculate with even numbers")
}
