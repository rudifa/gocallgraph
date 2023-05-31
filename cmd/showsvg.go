/*
Copyright © 2023 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/rudifa/gocallgraph/pkg/util"
)

// showitCmd represents the showit command
var showitCmd = &cobra.Command{
	Use:     "show",
	Short:   "Display the svg file in browser",
	Long:    `Display the svg file ` + outputsvgfile + `in browser.`,
	Aliases: []string{"showsvg"},
	Run: func(cmd *cobra.Command, args []string) {

		util.ShowSvg(outputsvgfile)
	},
}

func init() {
	rootCmd.AddCommand(showitCmd)
}
