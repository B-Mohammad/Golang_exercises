package main

import (
	"fmt"
	"os"
)

func main() {
	var sum float64
	var count int

	for {
		var temp float64
		_, err := fmt.Fscanln(os.Stdin, &temp)
		if err != nil {
			break
		}
		sum += temp
		count++
	}

	if count == 0 {
		fmt.Fprintln(os.Stderr, "No input")
		os.Exit(-1)
	}

	fmt.Fprintln(os.Stdout, "The AVG is:", sum/float64(count))
}
