package util

import (
	"testing"
	// "github.com/rudifa/gocallgraph/util"
)

// func TestStringList(t *testing.T) {
// 	sl := New()

// 	// should be empty
// 	if len(sl.List()) != 0 {
// 		t.Errorf("New() returned non-empty list")
// 	}

// 	// add some items
// 	sl.Add("a")
// 	sl.Add("b")
// 	sl.Add("c")

// 	// should have 3 items
// 	if len(sl.List()) != 3 {
// 		t.Errorf("List() returned wrong number of items")
// 	}

// 	// should contain "a"
// 	if !sl.List().contain("a") {
// 		t.Errorf("List() does not contain 'a'")
// 	}

// }

// test that List() returns a string
func TestStringList(t *testing.T) {
	sl := NewStringList()

	// should be empty
	if len(sl.List()) != 0 {
		t.Errorf("New() returned non-empty list")
	}

	// add some items
	sl.Add("abra")
	sl.Add("b")
	sl.Add("cadabra")

	// should have 3 items
	if len(sl.List()) != 3 {
		t.Errorf("List() returned wrong number of items")
	}

	// should contain "abra"
	if !sl.Contains("abra") {
		t.Errorf("List() does not contain 'a'")
	}

	// test AllMatching
	matches := sl.AllMatching("abra")
	if len(matches) != 2 {
		t.Errorf("AllMatching() returned wrong number of items")
	}

	// test FirstMatching
	match := sl.FirstMatching("abra")
	if match != "abra" {
		t.Errorf("FirstMatching() returned wrong item")
	}
}
