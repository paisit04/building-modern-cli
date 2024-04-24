/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// commandCmd represents the command command
var commandCmd = &cobra.Command{
	Use:   "command",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("command called")
		localFlag, _ := cmd.Flags().GetString("localFlag")
		if localFlag != "" {
			fmt.Printf("localFlag is set to %s\n", localFlag)
		}
	},
}

func init() {
	rootCmd.AddCommand(commandCmd)

	commandCmd.Flags().String("localFlag", "", "a local string flag")
	commandCmd.PersistentFlags().Bool("persistentFlag", false, "a persistent boolean flag")
}
