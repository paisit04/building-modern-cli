/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/paisit04/building-modern-cli/calculator/storage"
	"github.com/spf13/cobra"
)

// divideCmd represents the divide command
var divideCmd = &cobra.Command{
	Use:   "divide number",
	Short: "Divide value",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("only accepts a single argument")
			return
		}
		if len(args) == 0 {
			fmt.Println("command requires input value")
			return
		}
		floatVal, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Printf("unable to parse input[%s]: %v", args[0], err)
			return
		}
		value := storage.GetValue()
		value /= floatVal
		storage.SetValue(value)
		fmt.Printf("%f\n", value)
	},
}

func init() {
	rootCmd.AddCommand(divideCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// divideCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// divideCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
