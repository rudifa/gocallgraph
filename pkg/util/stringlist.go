/*
Copyright © 2023 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/

package util

import (
	"bufio"
	"errors"
	"os"
	"regexp"
)

// StringList wraps a string slice and provides methods to manipulate it
type StringList struct {
	list []string
}

// NewStringList returns an instance of empty StringList
func NewStringList() *StringList {
	return &StringList{}
}

// NewStringListFromFile returns an instance of StringList initialized from a file
func NewStringListFromFile(filename string) (*StringList, error) {
	sl := &StringList{}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sl.list = append(sl.list, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return sl, nil
}

// NewStringListFromSlice returns an instance of StringList initialized from a slice
func NewStringListFromSlice(slice []string) (*StringList, error) {
	if slice == nil {
		return nil, errors.New("input slice cannot be nil")
	}
	sl := &StringList{
		list: slice,
	}
	return sl, nil
}

// Add adds a string to the list
func (c *StringList) Add(call string) {
	c.list = append(c.list, call)
}

// List returns the list
func (c *StringList) List() []string {
	return c.list
}

// Contains returns true if the list contains the string
func (c *StringList) Contains(s string) bool {
	for _, v := range c.list {
		if v == s {
			return true
		}
	}
	return false
}

// AllMatching returns a list of strings matching the pattern
func (c *StringList) AllMatching(pattern string) []string {
	matches := make([]string, 0)
	for _, v := range c.list {
		if Matches(pattern, v) {
			matches = append(matches, v)
		}
	}
	return matches
}

// FirstMatching returns the first string matching the pattern
func (c *StringList) FirstMatching(pattern string) string {
	for _, v := range c.list {
		if Matches(pattern, v) {
			return v
		}
	}
	return ""
}

// should go into strings util

// Matches returns true if the string matches the pattern
func Matches(patterrn, s string) bool {
	match, _ := regexp.MatchString(patterrn, s)
	return match
}
