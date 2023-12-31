/*
Copyright © 2023 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/

package cmd

import (
	"log"
	"os"

	"github.com/rudifa/gocallgraph/pkg/util"
	"github.com/spf13/cobra"
)

const tmpdir = ".tmp"
const callgraphrawfile = ".tmp/callgraph.raw"
const callersfile = ".tmp/callers.txt"
const targetsfile = ".tmp/targets.txt"
const outputdotfile = ".tmp/callgraph.dot"
const outputsvgfile = ".tmp/callgraph.svg"

// CmdVerbose enables verbose output
var CmdVerbose bool = false

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gocallgraph",
	Short: "Gets call graph of a go module, converts it to dot and displays it",
	Long: `Gets call graph of a go module, converts it to dot and displays it...

	You might need to install graphviz: https://graphviz.org/download/
	and callgraph: https://pkg.go.dev/golang.org/x/tools/go/callgraph

	For brevity you can create a local alias
	  alias gcg=gocallgraph
	  alias gcg='go run'
	as appropriate.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if CmdVerbose {
			// propagate CmdVerbose flag to util
			util.Verbose = true
			log.Println("Verbose mode is enabled")
		}
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.EnablePrefixMatching = true                  // allow abbreviations
	rootCmd.CompletionOptions.DisableDefaultCmd = true // disable default completion

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&CmdVerbose, "verbose", "v", false, "Enable verbose output")
}
