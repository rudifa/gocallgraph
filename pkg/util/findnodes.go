/*
Copyright © 2023 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/
package util

func FindNodes(callgrapgfile string, matching string) ([]string, error) {
	calllist, err := NewCallListFromFile(callgrapgfile)
	if err != nil {
		return nil, err
	}

	matched := calllist.AllMatching(matching)
	return matched, nil
}
