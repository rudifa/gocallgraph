/*
Copyright © 2023 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/

package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/rudifa/gocallgraph/pkg/util"
)

// dotitCmd represents the showit command
var dotitCmd = &cobra.Command{
	Use:   "dot",
	Short: "Convert the raw call graph to dot file and to svg file",
	Long: `Convert a subset of the raw call graph to a grephviz dot file (` + outputdotfile + `).
	Also convert the dot file to a svg file (` + outputsvgfile + `) and display it in browser.

	The subset of the raw call graph is defined by the files ` + callersfile + ` and ` + targetsfile + `.

	You should look up callers and targets of interest in the file ./tmp/callgraph.raw and
	copy them to the file ./tmp/callers.txt.

	Be careful to copy the full function name, including the package name,
	e.g. 'github.com/rudifa/gocallgraph/cmd.Execute'.

	Display the svg file with show command, with LiveServer or with 'open .tmp/callgraph.svg'.`,
	Aliases: []string{"dotsvg"},
	Args:    cobra.NoArgs, // Ensure that no positional arguments are found.
	Run: func(cmd *cobra.Command, args []string) {

		horizontal, _ := cmd.Flags().GetBool("left-to-right")
		showsvg, _ := cmd.Flags().GetBool("show")

		err := util.RawToDot(callgraphrawfile, callersfile, targetsfile, horizontal, outputdotfile)
		if err != nil {
			log.Printf("0 error: %v", err)
			return
		}

		err = util.DotToSvg(outputdotfile, outputsvgfile)

		if err != nil {
			log.Printf("1 error: %v", err)
			return
		}

		if showsvg {
			util.ShowSvg(outputsvgfile)
		}

	},
}

func init() {
	rootCmd.AddCommand(dotitCmd)

	dotitCmd.Flags().BoolP("left-to-right", "l", false, "graph orientation left-to-right (default top-to-bottom)")
	dotitCmd.Flags().BoolP("show", "s", false, "show the svg file in browser")

}
