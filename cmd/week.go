/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"phoebe/storage"

	"github.com/spf13/cobra"
)

// weekCmd represents the week command
var weekCmd = &cobra.Command{
	Use:   "week",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		events, err := storage.LoadEvents(tomlPath)
		if err != nil {
			fmt.Printf("Your events could not be fetched: %v", err)
			os.Exit(1)
		}
		fmt.Printf("%s", events)
	},
}

func init() {
	rootCmd.AddCommand(weekCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// weekCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// weekCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}