/*
Copyright © 2023 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/rudifa/gocallgraph/pkg/util"
)

// dotitCmd represents the showit command
var dotitCmd = &cobra.Command{
	Use:   "svg",
	Short: "Convert the raw call graph to dot format and to svg file",
	Long: `Convert the raw call graph to dot file (` + outputdotfile + `) and to a svg file (` + outputsvgfile + `),
	for callers listed in the ` + callgraphrawfile + ` file.

	You should look up callers of interest to you in the file ./tmp/callgraph.raw and
	copy them to the file ./tmp/callers.txt.

	Be careful to copy the full function name, including the package name,
	e.g. 'github.com/rudifa/gocallgraph/cmd.Execute'.

	View the svg file with show command, with LiveServer or with 'open .tmp/callgraph.svg'.`,
	Aliases: []string{"dotsvg"},
	Run: func(cmd *cobra.Command, args []string) {

		util.RawToDot(callgraphrawfile, callersfile, calleesfile, outputdotfile)

		util.DotToSvg(outputdotfile, outputsvgfile)
	},
}

func init() {
	rootCmd.AddCommand(dotitCmd)
}
