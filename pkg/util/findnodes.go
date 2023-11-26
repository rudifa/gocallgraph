/*
Copyright © 2023 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/
package util

import "regexp"

func FindNodes(callgrapgfile string, matching string) ([]string, error) {
	nodes, err := extractNodes(callgrapgfile)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

// // DotToSvgAndShow converts the dot file to svg
// func DotToSvg(dotfile, svgfile string) {

// 	// Run the dot command
// 	stdout, stderr, err := fexec.RunCommand("dot", "-Tsvg", dotfile)

// 	if err != nil || stderr != "" {
// 		log.Fatal("*** DotToSvgAndShow stderr:", stderr)
// 	}

// 	// Write the svg to svgfile
// 	err = files.WriteToFile(svgfile, stdout)
// 	if err != nil {
// 		log.Fatalf("Failed to write SVG file: %v", err)
// 	}

// 	log.Println("Wrote SVG file:", svgfile)
// }

// // ShowSvg opens the svg file in browser
// func ShowSvg(svgfile string) {

// 	// Open the svg file in browser
// 	_, _, err := fexec.RunCommand("open", svgfile)
// 	if err != nil {
// 		log.Fatalf("Failed to open SVG file: %v", err)
// 	}
// }

func matches(s string, matching string) bool {
	match, _ := regexp.MatchString(matching, s)
	return match
}
