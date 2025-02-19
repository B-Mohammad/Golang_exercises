package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// for _, fileName := range os.Args[1:] {

	// 	file, err := os.Open(fileName)
	// 	if err != nil {
	// 		fmt.Fprintln(os.Stderr, err.Error())
	// 		continue
	// 	}

	// 	if _, err := io.Copy(os.Stdout, file); err != nil {
	// 		fmt.Fprintln(os.Stderr, err.Error())
	// 		continue
	// 	}

	// 	file.Close()
	// }

	// for _, fileName := range os.Args[1:] {

	// 	file, err := os.Open(fileName)
	// 	if err != nil {
	// 		fmt.Fprintln(os.Stderr, err.Error())
	// 		continue
	// 	}

	// 	data, err := io.ReadAll(file)
	// 	if err != nil {
	// 		fmt.Fprintln(os.Stderr, err.Error())
	// 		continue
	// 	}

	// 	fmt.Println("The file has", len(data), "bytes")

	// 	file.Close()
	// }
	var tcc, tcw, tcl int

	for _, fileName := range os.Args[1:] {
		var cc, cw, cl int

		file, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			continue
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			cl++
			cc += len(line)
			cw += len(strings.Fields(line))
		}

		tcc += cc
		tcw += cw
		tcl += cl

		fmt.Printf("File %-12s: %6d bytes, %6d words, %6d lines\n", fileName, cc, cw, cl)

		file.Close()
	}

	fmt.Printf("File %-12s: %6d bytes, %6d words, %6d lines\n", "Total", tcc, tcw, tcl)

}
