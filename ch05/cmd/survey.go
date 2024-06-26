/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// surveyCmd represents the survey command
var surveyCmd = &cobra.Command{
	Use:   "survey",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("survey called")

		var qs = []*survey.Question{
			{
				Name:      "firstname",
				Prompt:    &survey.Input{Message: "What is your first name?"},
				Validate:  survey.Required,
				Transform: survey.Title,
			},
			{
				Name: "favoritecolor",
				Prompt: &survey.Select{
					Message: "What's your favorite color?",
					Options: []string{"red", "orange", "yellow", "green", "blue", "purple", "black", "brown", "white"},
					Default: "white",
				},
			},
			{
				Name: "story",
				Prompt: &survey.Multiline{
					Message: "Tell me a story.",
				},
			},
			{
				Name: "secret",
				Prompt: &survey.Password{
					Message: "Tell me a secret",
				},
			},
			{
				Name: "good",
				Prompt: &survey.Confirm{
					Message: "Are you having a good day?",
				},
			},
			{
				Name: "favoritepies",
				Prompt: &survey.MultiSelect{
					Message: "What pies do you like:",
					Options: []string{"Pumpkin", "Lemon Meringue", "Cherry", "Apple", "Key Lime", "Pecan", "Boston Cream", "Rhubarb", "Blackberry"},
				},
			},
		}

		answers := struct {
			FirstName     string
			FavoriteColor string
			Story         string
			Secret        string
			Good          bool
			FavoritePies  []string
		}{}
		err := survey.Ask(qs, &answers)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("*********** SURVEY RESULTS ***********")
		fmt.Printf("First Name: %s\n", answers.FirstName)
		fmt.Printf("Favorite Color: %s\n", answers.FavoriteColor)
		fmt.Printf("Story: %s\n", answers.Story)
		fmt.Printf("Secret: %s\n", answers.Secret)
		fmt.Printf("It's a good day: %v\n", answers.Good)
		fmt.Printf("Favorite Pies: %s\n", answers.FavoritePies)
	},
}

func init() {
	rootCmd.AddCommand(surveyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// surveyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// surveyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
