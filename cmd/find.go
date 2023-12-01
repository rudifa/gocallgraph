/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"log"
	"regexp"

	"github.com/rudifa/gocallgraph/pkg/util"
	"github.com/spf13/cobra"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find a function in the raw callgraph",
	Long: `Find a function in the raw callgraph
that matches the given function signature string.`,
	Args: cobra.NoArgs, // Ensure that no positional arguments are found.
	Run: func(cmd *cobra.Command, args []string) {
		funcsig, _ := cmd.Flags().GetString("function-signature")

		found, err := FindFunctionInRawCallgraph(funcsig, callgraphrawfile)

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

func init() {
	rootCmd.AddCommand(findCmd)

	findCmd.Flags().StringP("function-signature", "f", "", "function signature to match")

	findCmd.MarkFlagRequired("function-signature")
}

// FindFunctionInRawCallgraph finds functions in the raw callgraph
// that match the funcsig.
func FindFunctionInRawCallgraph(funcsig string, filename string) ([]string, error) {
	pattern := regexp.QuoteMeta(funcsig)
	if CmdVerbose {
		log.Println("function-signature:", funcsig)
		fmt.Println("pattern:", pattern)
	}
	matched, err := util.FindMatchingNodes(filename, pattern)

	return matched, err
}
