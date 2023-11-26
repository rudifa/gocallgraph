/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find a function in the raw callgraph",
	Long: `Find a function in the raw callgraph
that matches the given string.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("find called", regex)
	},
}

var regex string

func init() {
	rootCmd.AddCommand(findCmd)

	// add -r --regex option
	findCmd.Flags().StringVarP(&regex, "regex", "r", "", "regex pattern to match")

	findCmd.MarkFlagRequired("regex")
}

// FindFunctionInRawCallgraph finds a function in the raw callgraph
// that matches the given string.
func FindFunctionInRawCallgraph(regex string, filename string) {
	fmt.Println("FindFunctionInRawCallgraph called", regex, filename)
}
