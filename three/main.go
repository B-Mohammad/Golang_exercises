package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Not enough arguments")
		os.Exit(-1)
	}

	oldV, newV := os.Args[1], os.Args[2]

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		temp := strings.Split(scanner.Text(), oldV)
		result := strings.Join(temp, newV)

		fmt.Fprintln(os.Stdout, result)

	}

}
