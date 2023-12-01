/*
Copyright © 2023 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/
package util

import (
	"errors"
	"log"

	"github.com/rudifa/goutil/fexec"
	"github.com/rudifa/goutil/files"
)

// DotToSvg converts the dot file to svg
func DotToSvg(dotfile, svgfile string) error {

	// Run the dot command
	stdout, stderr, err := fexec.RunCommand("dot", "-Tsvg", dotfile)

	if err != nil {
		return err
	}

	if stderr != "" {
		err := errors.New(stderr)
		return err
	}

	// Write the svg to svgfile
	err = files.WriteToFile(svgfile, stdout)
	if err != nil {
		return err
	}

	log.Println("Wrote SVG file:", svgfile)

	return nil
}

// ShowSvg opens the svg file in browser
func ShowSvg(svgfile string) {

	// Open the svg file in browser
	_, _, err := fexec.RunCommand("open", svgfile)
	if err != nil {
		log.Fatalf("Failed to open SVG file: %v", err)
	}
}
