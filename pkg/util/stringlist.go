package util

import (
	"bufio"
	"errors"
	"log"
	"os"
)

type StringList struct {
	list []string
}

// init from a file
func (c *StringList) initFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// parse line and add to list
		// assuming Calllist has a method Add and a way to parse a line into a call
		// call := parseLineToCall(scanner.Text())
		// c.Add(call)
		c.Add(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (c *StringList) Add(call string) {
	c.list = append(c.list, call)
}

func New() *StringList {
	return &StringList{}
}

func NewFromFile(filename string) (*StringList, error) {
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

func NewFrom(slice []string) (*StringList, error) {
	if slice == nil {
		return nil, errors.New("input slice cannot be nil")
	}
	sl := &StringList{
		list: slice,
	}
	return sl, nil
}
