/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"

	"github.com/rudifa/gocallgraph/pkg/util"
	"github.com/spf13/cobra"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find a function in the raw callgraph",
	Long: `Find a function in the raw callgraph
that matches the given string.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("find called with pattern:", pattern)

		//func FindMatchingNodes(callgrapgfile string, matching string) ([]string, error) {

		// found, err := util.FindMatchingNodes("testdata/callgraph.raw", pattern)

		found, err := FindFunctionInRawCallgraph(pattern, callgraphrawfile)

		if err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Println("found:", len(found))
		for _, f := range found {
			fmt.Println(f)
		}
	},
}

var pattern string

func init() {
	rootCmd.AddCommand(findCmd)

	// add -r --regex option
	findCmd.Flags().StringVarP(&pattern, "pattern", "p", "", "regex pattern to match")

	findCmd.MarkFlagRequired("pattern")
}

// FindFunctionInRawCallgraph finds functions in the raw callgraph
// that match the given pattern.
func FindFunctionInRawCallgraph(pattern string, filename string) ([]string, error) {
	// fmt.Println("FindFunctionInRawCallgraph called", pattern, filename)

	matched, err := util.FindMatchingNodes(filename, pattern)

	return matched, err
}
