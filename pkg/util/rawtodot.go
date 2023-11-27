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

func RawToDot(callgraphfile, callersfile, calleesfile, outputdotfile string) error {

	// Get callers of interest from the callers file
	callers, err := processCallersFile(callersfile)
	if err != nil {
		log.Fatal(err)
	}

	// print the callers
	log.Println("Callers:")
	for _, caller := range callers {
		log.Println(caller)
	}

	// Get callees of interest from the callees file
	callees, err := processCalleesFile(calleesfile)
	if err != nil {
		log.Fatal(err)
	}

	// print the callees
	log.Println("Callees:")
	for _, callee := range callees {
		log.Println(callee)
	}

	// Get all tuples (caller, callee) from the callgraph.raw file
	tuples, err := callgraphRawToTuples(callgraphfile)
	if err != nil {
		log.Fatal(err)
	}

	extractRootsAndLeavesToFiles(tuples)

	// Create the DOT file where nodes are callers
	// and edges defined by the corresponding tuples
	err = generateDotFile(tuples, callers, callees, outputdotfile)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// callgraphRawToTuples reads the calltree file and returns an array of tuples (caller, callee)
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

		caller, callee, err := parseLine(line)
		if err != nil {
			// ignore error
			continue
		}
		// append the tuple to the array
		tuples = append(tuples, [2]string{caller, callee})
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

// processCalleesFile reads the callers file and returna an array of callees
func processCalleesFile(filename string) ([]string, error) {

	callees := make([]string, 0)
	file, err := os.Open(filename)
	if err != nil {
		return callees, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return callees, err
		}
		callees = append(callees, strings.TrimSpace(line))
	}

	return callees, err
}

// extractRootsAndLeavesToFiles finds all the roots and all the leaves of the calltree
// and writes them to the roots.txt and leaves.txt
func extractRootsAndLeavesToFiles(tuples [][2]string) {

	// a root is a caller that is not a callee
	// a leaf is a callee that is not a caller

	roots := make([]string, 0)
	leaves := make([]string, 0)

	// should use maps instead of slices to extract roots and leaves

	// make maps of callers and callees to be used as sets

	callers := make(map[string]bool)
	callees := make(map[string]bool)

	for _, tuple := range tuples {
		callers[tuple[0]] = true
		callees[tuple[1]] = true
	}

	// loop over the callees and check if they are not in the callers
	for callee := range callees {
		if _, ok := callers[callee]; !ok {
			leaves = append(leaves, callee)
		}
	}

	// loop over the callers and check if they are not in the callees
	for caller := range callers {
		if _, ok := callees[caller]; !ok {
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
func generateDotFile(tuples [][2]string, callers []string, callees []string, outputdotfile string) error {
	dotFile, err := os.Create(outputdotfile)
	if err != nil {
		return err
	}
	defer dotFile.Close()

	writer := bufio.NewWriter(dotFile)
	defer writer.Flush()

	// write the header
	_, err = writer.WriteString("digraph GraphName {\n\trankdir=TB;\n")
	if err != nil {
		return err
	}

	// loop over callers and write edge lines for each caller
	for _, caller := range callers {
		_, err = writer.WriteString(fmt.Sprintf("\t\"%s\"[shape=box, style=\"rounded,filled\", fillcolor=\"yellow\", color=black];\n", caller))

		if err != nil {
			return err
		}
		callertuples := getTuplesForCaller(caller, tuples)
		// loop over the tuples and write the edge lines
		for _, tuple := range callertuples {
			_, err = writer.WriteString(fmt.Sprintf("\t\"%s\" -> \"%s\";\n", tuple[0], tuple[1]))
			if err != nil {
				return err
			}
		}
	}

	// loop over callees and write edge lines for each callee
	for _, callee := range callees {
		_, err = writer.WriteString(fmt.Sprintf("\t\"%s\"[shape=box, style=\"rounded,filled\", fillcolor=\"aquamarine\", color=black];\n", callee))

		if err != nil {
			return err
		}
		calleetuples := getTuplesForCallee(callee, tuples)
		// loop over the tuples and write the edge lines
		for _, tuple := range calleetuples {
			_, err = writer.WriteString(fmt.Sprintf("\t\"%s\" -> \"%s\";\n", tuple[0], tuple[1]))
			if err != nil {
				return err
			}
		}
	}

	// write the footer
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

// getTuplesForCallee returns an array of tuples where the callee is the second element of the tuple
func getTuplesForCallee(callee string, tuples [][2]string) [][2]string {
	result := make([][2]string, 0)
	for _, tuple := range tuples {
		if tuple[1] == callee {
			result = append(result, tuple)
		}
	}
	return result
}

// parseLine parses a line of the callgraph.raw file and returns the caller and callee
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

		caller, callee, err := parseLine(line)
		if err != nil {
			// ignore error
			continue
		}
		nodes.Insert(caller)
		nodes.Insert(callee)
	}

	list := make([]string, 0)
	return list, err
}
