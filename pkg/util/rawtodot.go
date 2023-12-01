/*
Copyright © 2023 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/
package util

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/golang-collections/collections/set"
)

// RawToDot converts the callgraph.raw file to a DOT file, filtering by callers and targets
func RawToDot(callgraphfile, callersfile, targetsfile string, horizontal bool, outputdotfile string) error {

	// Get callers of interest from the callers file
	callers, err := processCallersFile(callersfile)
	if err != nil {
		return err
	}

	// print the callers
	log.Println("callers:")
	for _, caller := range callers {
		log.Println(caller)
	}

	// Get targets of interest from the targets file
	targets, err := processTargetsFile(targetsfile)
	if err != nil {
		return err
	}

	// print the targets
	log.Println("targets:")
	for _, target := range targets {
		log.Println(target)
	}

	// Get all tuples (caller, target) from the callgraph.raw file
	tuples, err := callgraphRawToTuples(callgraphfile)
	if err != nil {
		return err
	}

	extractRootsAndLeavesToFiles(tuples)

	// Create the DOT file where nodes are callers
	// and edges defined by the corresponding tuples
	err = generateDotFile(tuples, callers, targets, horizontal, outputdotfile)
	if err != nil {
		return err
	}

	return nil
}

// callgraphRawToTuples reads the calltree file and returns an array of tuples (caller, target)
func callgraphRawToTuples(filename string) ([][2]string, error) {
	tuples := make([][2]string, 0)

	file, err := os.Open(filename)
	if err != nil {
		return tuples, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return tuples, err
		}

		caller, target, err := parseLine(line)
		if err != nil {
			// ignore error
			continue
		}
		// append the tuple to the array
		tuples = append(tuples, [2]string{caller, target})
	}

	return tuples, err
}

// processCallersFile reads the callers file and returns an array of callers
func processCallersFile(filename string) ([]string, error) {
	callers := make([]string, 0)
	file, err := os.Open(filename)
	if err != nil {
		return callers, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return callers, err
		}
		callers = append(callers, strings.TrimSpace(line))
	}

	return callers, err
}

// processTargetsFile reads the callers file and returna an array of targets
func processTargetsFile(filename string) ([]string, error) {

	targets := make([]string, 0)
	file, err := os.Open(filename)
	if err != nil {
		return targets, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return targets, err
		}
		targets = append(targets, strings.TrimSpace(line))
	}

	return targets, err
}

// extractRootsAndLeavesToFiles finds all the roots and all the leaves of the calltree
// and writes them to the roots.txt and leaves.txt
func extractRootsAndLeavesToFiles(tuples [][2]string) {

	// a root is a caller that is not a target
	// a leaf is a target that is not a caller

	roots := make([]string, 0)
	leaves := make([]string, 0)

	// should use maps instead of slices to extract roots and leaves

	// make maps of callers and targets to be used as sets

	callers := make(map[string]bool)
	targets := make(map[string]bool)

	for _, tuple := range tuples {
		callers[tuple[0]] = true
		targets[tuple[1]] = true
	}

	// loop over the targets and check if they are not in the callers
	for target := range targets {
		if _, ok := callers[target]; !ok {
			leaves = append(leaves, target)
		}
	}

	// loop over the callers and check if they are not in the targets
	for caller := range callers {
		if _, ok := targets[caller]; !ok {
			roots = append(roots, caller)
		}
	}

	// sort leaves and roots
	sort.Strings(leaves)
	sort.Strings(roots)

	// write the roots to the roots.txt file

	rootsFile, err := os.Create(".tmp/roots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer rootsFile.Close()

	rootsWriter := bufio.NewWriter(rootsFile)
	defer rootsWriter.Flush()

	for _, root := range roots {
		_, err := rootsWriter.WriteString(root + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	if Verbose {
		log.Println("Wrote roots to .tmp/roots.txt")
	}

	// write the leaves to the leaves.txt file

	leavesFile, err := os.Create(".tmp/leaves.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer leavesFile.Close()

	leavesWriter := bufio.NewWriter(leavesFile)
	defer leavesWriter.Flush()

	for _, leaf := range leaves {
		_, err := leavesWriter.WriteString(leaf + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	if Verbose {
		log.Println("Wrote leaves to .tmp/leaves.txt")
	}

}

// generateDotFile writes the outputdotfile with the data from the tuples and callers
func generateDotFile(tuples [][2]string, callers, targets []string, horizontal bool, outputdotfile string) error {
	dotFile, err := os.Create(outputdotfile)
	if err != nil {
		return err
	}
	defer dotFile.Close()

	writer := bufio.NewWriter(dotFile)
	defer writer.Flush()

	// write the dot dotHeader
	rankdir := "TB"
	if horizontal {
		rankdir = "LR"
	}
	dotHeader := fmt.Sprintf("digraph GraphName {\n\trankdir=%s;\n", rankdir)
	_, err = writer.WriteString(dotHeader)
	if err != nil {
		return err
	}

	done := set.New()

	// loop over callers and write edge lines for each caller
	for _, caller := range callers {
		if caller == "" {
			continue
		}
		_, err = writer.WriteString(fmt.Sprintf("\t\"%s\"[shape=box, style=\"rounded,filled\", fillcolor=\"yellow\", color=black];\n", caller))

		if err != nil {
			return err
		}
		callertuples := getTuplesForCaller(caller, tuples)
		// loop over the tuples and write the edge lines
		for _, tuple := range callertuples {
			if done.Has(tuple) {
				continue
			}
			done.Insert(tuple)

			_, err = writer.WriteString(fmt.Sprintf("\t\"%s\" -> \"%s\";\n", tuple[0], tuple[1]))
			if err != nil {
				return err
			}
		}
	}

	// loop over targets and write edge lines for each target
	for _, target := range targets {
		if target == "" {
			continue
		}

		_, err = writer.WriteString(fmt.Sprintf("\t\"%s\"[shape=box, style=\"rounded,filled\", fillcolor=\"aquamarine\", color=black];\n", target))

		if err != nil {
			return err
		}
		targettuples := getTuplesForTarget(target, tuples)

		// loop over the tuples and write the edge lines
		for _, tuple := range targettuples {
			if done.Has(tuple) {
				continue
			}
			done.Insert(tuple)

			_, err = writer.WriteString(fmt.Sprintf("\t\"%s\" -> \"%s\";\n", tuple[0], tuple[1]))
			if err != nil {
				return err
			}
		}
	}

	// write the dot footer
	_, err = writer.WriteString("}")
	if err != nil {
		return err
	}

	log.Println("DOT file generated successfully:", dotFile.Name())
	return nil
}

// getTuplesForCaller returns an array of tuples where the caller is the first element of the tuple
func getTuplesForCaller(caller string, tuples [][2]string) [][2]string {
	result := make([][2]string, 0)
	for _, tuple := range tuples {
		if tuple[0] == caller {
			result = append(result, tuple)
		}
	}
	return result
}

// getTuplesForTarget returns an array of tuples where the target is the second element of the tuple
func getTuplesForTarget(target string, tuples [][2]string) [][2]string {
	result := make([][2]string, 0)
	for _, tuple := range tuples {
		if tuple[1] == target {
			result = append(result, tuple)
		}
	}
	return result
}

// parseLine parses a line of the callgraph.raw file and returns the caller and target
func parseLine(line string) (string, string, error) {
	fragments := strings.Split(strings.TrimSpace(line), "\t")
	if len(fragments) != 3 {
		return "", "", fmt.Errorf("invalid line format")
	}
	return fragments[0], fragments[2], nil
}

func extractNodes(filename string) ([]string, error) {
	nodes := set.New()
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			// ignore error
			continue
		}

		caller, target, err := parseLine(line)
		if err != nil {
			// ignore error
			continue
		}
		nodes.Insert(caller)
		nodes.Insert(target)
	}

	list := make([]string, 0)
	return list, err
}
