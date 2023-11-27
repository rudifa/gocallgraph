/*
Copyright © 2023 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/

package util

import (
	"log"
	"sort"
	"strings"

	"github.com/rudifa/goutil/fexec"
	"github.com/rudifa/goutil/files"
)

// GetRawCallgraph runs the callgraph command and writes the output to the filepath
func GetRawCallgraph(filepath string) {

	// Run the callgraph command
	stdout, stderr, err := fexec.RunCommand("callgraph", ".")

	if stderr != "" {
		// stderr = util.Truncate(stderr, 200)
		log.Fatal("*** GetRawCallgraph stderr:", stderr)
	}

	if err != nil {
		log.Fatal("*** GetRawCallgraph error:", err)
	}

	lines := strings.Split(stdout, "\n")

	sort.Strings(lines)

	joined := strings.Join(lines, "\n")

	// write to file
	err = files.WriteToFile(filepath, joined)
	if err != nil {
		log.Fatal("*** GetRawCallgraph error:", err)
	}

	log.Printf("Wrote the raw callgraph file %v", filepath)
}
