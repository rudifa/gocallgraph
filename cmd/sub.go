/*
/*
Copyright © 2023 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// subCmd represents the sub command
var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "Remove a function to visualize, as a caller or as a target.",
	Long:  "Remove a function to visualize, as a caller or as a target.",
	Args:  cobra.NoArgs, // Ensure that no positional arguments are found.
	Run: func(cmd *cobra.Command, args []string) {
		caller, _ := cmd.Flags().GetString("caller")
		target, _ := cmd.Flags().GetString("target")
		fmt.Println("sub caller:", caller)
		fmt.Println("sub target:", target)

		if caller != "" {
			RemoveFrom(callersfile, caller)
		}
		if target != "" {
			RemoveFrom(targetsfile, target)
		}
	},
}

func init() {
	rootCmd.AddCommand(subCmd)

	subCmd.Flags().StringP("caller", "c", "", "function to add as caller")
	subCmd.Flags().StringP("target", "t", "", "function to add as target")

	// subCmd.MarkFlagRequired("caller")
	// subCmd.MarkFlagRequired("target")
}
