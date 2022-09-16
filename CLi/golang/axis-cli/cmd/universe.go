/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"

	"axis-cli/cmd/universe"
)

// universeCmd represents the universe command
var universeCmd = &cobra.Command{
	Use:   "universe",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func init() {
	universeCmd.AddCommand(universe.CreateCmd)
	universeCmd.AddCommand(universe.ListCmd)
	universeCmd.AddCommand(universe.ShowCmd)

	rootCmd.AddCommand(universeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// itemCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// itemCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}