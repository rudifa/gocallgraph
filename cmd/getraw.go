/*
Copyright © 2023 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/rudifa/gocallgraph/pkg/util"
	"github.com/rudifa/goutil/files"
)

// getcallgraphCmd represents the getraw command
var getcallgraphCmd = &cobra.Command{
	Use:     "raw",
	Short:   "Get the raw call for the currrent module",
	Long:    `Get the raw call for the currrent module...`,
	Aliases: []string{"getraw"},
	Run: func(cmd *cobra.Command, args []string) {

		files.EnsureDirectoryExists(tmpdir)

		util.GetRawCallgraph(callgraphrawfile)
	},
}

func init() {
	rootCmd.AddCommand(getcallgraphCmd)
}
