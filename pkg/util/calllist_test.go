package util

import (
	"log"
	"strings"
	"testing"
)

func TestCallList(t *testing.T) {
	cl := NewCallList()

	// should be empty
	if len(cl.List()) != 0 {
		t.Errorf("New() returned non-empty list")
	}

	// add some items
	cl.Add("a")
	cl.Add("b")
	cl.Add("c")

	// should have 3 items
	if len(cl.List()) != 3 {
		t.Errorf("List() returned wrong number of items")
	}

	// should contain "a"
	if !cl.Contains("a") {
		t.Errorf("List() does not contain 'a'")
	}

}

// test init from strings
func TestCallListFromStrings(t *testing.T) {

	const input = `github.com/rudifa/gocallgraph.init	--static-0:0-->	github.com/rudifa/gocallgraph/cmd.init
github.com/rudifa/gocallgraph.main	--static-9:13-->	github.com/rudifa/gocallgraph/cmd.Execute
github.com/rudifa/gocallgraph/cmd.Execute	--static-49:24-->	(*github.com/spf13/cobra.Command).Execute
github.com/rudifa/gocallgraph/cmd.Execute	--static-51:10-->	os.Exit`

	// split input on newlines
	strings := strings.Split(input, "\n")

	log.Println("len(strings) = ", len(strings))

	// if len(strings) != 4 {
	// 	t.Errorf("wrong number of lines")
	// }

	// log.Println("strings = ", strings)

	// print lines in a loop
	// for _, line := range strings {
	// 	log.Println("line = ", line)
	// }

	// create a new CallList
	cl := NewCallListFromStrings(strings)

	// log.Printf("cl.List() = %v", cl.List())

	// should have 4 items
	if len(cl.List()) != 4 {
		t.Errorf("List() returned wrong number of items")
	}

	// print List lines in a loop
	for _, line := range cl.List() {
		log.Println("line = ", line)
	}

	// get list of Nodes and check length
	nodes := cl.Nodes()

	// print Nodes in a loop
	for _, node := range nodes.List() {
		log.Println("node = ", node)
	}

	if len(nodes.List()) != 6 {
		t.Errorf("Nodes() returned wrong number of items")
	}

}

func TestCallListFromFile(t *testing.T) {
	// create a new CallList
	cl, err := NewCallListFromFile("testdata/callgraph.raw")
	if err != nil {
		t.Errorf("NewCallListFromFile() returned error: %v", err)
	}

	// should have 4 items
	if len(cl.List()) != 4 {
		t.Errorf("List() returned wrong number of items")
	}

	// print List lines in a loop
	for _, line := range cl.List() {
		log.Println("line = ", line)
	}

	// get list of Nodes and check length
	nodes := cl.Nodes()

	// print Nodes in a loop
	for _, node := range nodes.List() {
		log.Println("node = ", node)
	}

	if len(nodes.List()) != 6 {
		t.Errorf("Nodes() returned wrong number of items")
	}

	matching := nodes.AllMatching("init")

	log.Println("matching = ", matching)

	if len(matching) != 2 {
		t.Errorf("AllMatching() returned wrong number of items")
	}

}
