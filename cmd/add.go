/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

package cmd

import (
	"fmt"
	"strings"

	"os"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a function to visualize, as a caller or as a target.",
	Long:  `Add a function to visualize, as a caller or as a target of calls.`,
	Args:  cobra.NoArgs, // Ensure that no positional arguments are found.
	Run: func(cmd *cobra.Command, args []string) {
		caller, _ := cmd.Flags().GetString("caller")
		target, _ := cmd.Flags().GetString("target")
		fmt.Println("add caller:", caller)
		fmt.Println("add target:", target)

		if caller != "" {
			AppendTo(callersfile, caller)
		}
		if target != "" {
			AppendTo(targetsfile, target)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("caller", "c", "", "function to add as caller")
	addCmd.Flags().StringP("target", "t", "", "function to add as target")

	// addCmd.MarkFlagRequired("caller")
	// addCmd.MarkFlagRequired("target")
}

// AppendTo appends text to the file
func AppendTo(filename, text string) error {

	// Open the file in append mode. If the file doesn't exist, create it.
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening or creating file:", err)
		return err
	}
	defer file.Close()

	// Write the line string and a newline to the file.
	_, err = fmt.Fprintln(file, text)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	return nil
}

// RemoveFrom removes text from the file
func RemoveFrom(filename, text string) error {
	// Read the file
	contentBytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	// Convert to string and replace the text
	content := string(contentBytes)
	content = strings.Replace(content, text, "", -1)

	// Write the modified content back to the file
	err = os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}

	return nil
}
