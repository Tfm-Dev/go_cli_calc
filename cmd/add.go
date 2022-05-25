package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Addition calculation",
	Long: `Addition calculation with numbers integer or floating.
Need two or more arguments to be added.`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		float, _ := cmd.Flags().GetBool("float")
		if float {
			addFloat(args)
		} else {
			odd, _ := cmd.Flags().GetBool("odd")
			even, _ := cmd.Flags().GetBool("even")
			sumInt(args, odd, even)
		}
	},
}

func sumInt(args []string, odd bool, even bool) {
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
	var add float64
	for _, value := range args {
		temp, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Println(err)
		}
		add = add + temp
	}
	fmt.Printf("add of numbers: %s is %f\n", args, add)
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("float", "f", false, "add floating numbers")
	addCmd.Flags().BoolP("odd", "o", false, "calculate with odd numbers")
	addCmd.Flags().BoolP("even", "e", false, "calculate with even numbers")
}
