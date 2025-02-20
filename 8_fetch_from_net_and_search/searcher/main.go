package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type comicInfo struct {
	Link       string `json:"link"`
	Date       string `json:"date"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
}

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stdout, "not enough args")
		os.Exit(-1)
	}

	if len(os.Args) == 2 {
		fmt.Println("no query entry")
		os.Exit(0)
	}

	fp := os.Args[1]
	query := os.Args[2:]

	file, err := os.Open(fp)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can not open file", err.Error())
		os.Exit(-1)
	}

	defer file.Close()

	var comics []comicInfo
	if err := json.NewDecoder(file).Decode(&comics); err != nil {
		fmt.Fprintln(os.Stderr, "can not parse file", err.Error())
		os.Exit(-1)
	}

	var found int
	fmt.Printf("read %d comics\n", len(comics))

nextComic:
	for i := range comics {
		vti := strings.ToLower(comics[i].Title)
		vtc := strings.ToLower(comics[i].Transcript)
		for _, q := range query {
			q = strings.ToLower(q)

			if !strings.Contains(vti, q) && !strings.Contains(vtc, q) {
				continue nextComic
			}
		}
		fmt.Printf("%s %s %q\n", comics[i].Link, comics[i].Date, comics[i].Title)
		found++
	}
	fmt.Printf("found %d comic\n", found)

}
