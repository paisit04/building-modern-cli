/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// subcommandCmd represents the subcommand command
var subcommandCmd = &cobra.Command{
	Use:   "subcommand",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("subcommand called")
		} else {
			fmt.Println("subcommand called with args:", args)
		}
		persistentFlag, _ := cmd.Flags().GetBool("persistentFlag")
		fmt.Printf("persistentFlag is set to %v\n", persistentFlag)
	},
}

func init() {
	commandCmd.AddCommand(subcommandCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subcommandCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subcommandCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
