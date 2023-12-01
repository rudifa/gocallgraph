/*
Copyright © 2023 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/

package util

import (
	"errors"
	"strings"
)

// Verbose enables verbose output
var Verbose bool = false

// TODO add a helpers file

//  err := errors.New(stderr)

// FileError returns a customized error message
func FileError(err error, filename string) error {

	if strings.Contains(err.Error(), "no such file or directory") {
		err = errors.New("file not found: " + filename)
	}

	return err
}
