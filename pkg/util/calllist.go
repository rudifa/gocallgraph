package util

import (
	"strings"

	"github.com/golang-collections/collections/set"
)

// CallList contails a list of strings
// that contain space separated triplets (caller, call, callee)
type CallList struct {
	StringList // CallList is a StringList with more methods
}

// NewCallList returns an instance of CallList
func NewCallList() *CallList {
	return &CallList{}
}

// NewCallListFromFile returns an instance of CallList from a file
func NewCallListFromFile(filename string) (*CallList, error) {
	// use NewStringListFromFile to read the file
	stringList, err := NewStringListFromFile(filename)
	if err != nil {
		return nil, err
	}

	// create a new CallList
	callList := &CallList{
		StringList: *stringList,
	}

	return callList, nil
}

// NewCallListFromStrings returns an instance of CallList from a list of strings
func NewCallListFromStrings(strings []string) *CallList {

	stringlist, _ := NewStringListFromSlice(strings)
	callList := &CallList{
		StringList: *stringlist,
	}

	return callList
}

// Nodes returns a list of nodes (callers and callees)
func (c *CallList) Nodes() *StringList {
	// init a Set
	nodes := set.New()

	// for each line in List() split it into 3 parts on whitespace
	for _, line := range c.List() {
		parts := strings.Fields(line)
		if len(parts) >= 3 {
			// add caller and callee to the set
			nodes.Insert(parts[0])
			nodes.Insert(parts[2])
		}
	}

	// convert the set to a StringList
	list := NewStringList()
	nodes.Do(func(node interface{}) {
		list.Add(node.(string))
	})

	// return the set as a StringList
	return list
}
