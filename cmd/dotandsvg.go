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
	Use:   "dot",
	Short: "Convert the raw call graph to dot file and to svg file",
	Long: `Convert a subset of the raw call graph to a grephviz dot file (` + outputdotfile + `).
	Also convert the dot file to a svg file (` + outputsvgfile + `) and display it in browser.

	The subset of the raw call graph is defined by the files ` + callersfile + ` and ` + calleesfile + `.

	You should look up callers and callees of interest in the file ./tmp/callgraph.raw and
	copy them to the file ./tmp/callers.txt.

	Be careful to copy the full function name, including the package name,
	e.g. 'github.com/rudifa/gocallgraph/cmd.Execute'.

	Display the svg file with show command, with LiveServer or with 'open .tmp/callgraph.svg'.`,
	Aliases: []string{"dotsvg"},
	Run: func(cmd *cobra.Command, args []string) {

		util.RawToDot(callgraphrawfile, callersfile, calleesfile, outputdotfile)

		util.DotToSvg(outputdotfile, outputsvgfile)
	},
}

func init() {
	rootCmd.AddCommand(dotitCmd)
}
