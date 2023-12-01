/*
Copyright © 2023 Rudolf Farkas @rudifa rudi.farkas@gmail.com
*/
package main

import (
	"github.com/rudifa/gocallgraph/cmd"
)

func main() {
	cmd.Execute()
	// fib := fibonacci(5)
	// fmt.Println(fib)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
