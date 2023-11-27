/*
Copyright © 2023 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/

package util

// FindMatchingNodes finds all nodes in the callgraph that match the given pattern.
func FindMatchingNodes(callgrapgfile string, pattern string) ([]string, error) {
	calllist, err := NewCallListFromFile(callgrapgfile)
	if err != nil {
		return nil, err
	}

	nodes := calllist.Nodes()

	matched := nodes.AllMatching(pattern)
	return matched, nil
}
