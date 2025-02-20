package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type comicInfo struct {
	Link       string `json:"link"`
	Date       string `json:"date"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
}

type dataReceived struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Year       string `json:"year"`
	Title      string `json:"title"`
	Day        string `json:"day"`
	Transcript string `json:"transcript"`
}

const baseUrl = "https://xkcd.com/"

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Please insert dest file arg for record. Usage: go run main.go <dest>")
		os.Exit(-1)
	}

	var failed int
	comics := make([]comicInfo, 0, 3100)

	i := 1
	for failed < 2 && i < 1000 {

		res, err := http.Get(baseUrl + fmt.Sprint(i) + "/info.0.json")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching URL comic num %d\n", i)
			i++
			continue
		}

		if res.StatusCode == http.StatusNotFound {
			failed++
			fmt.Fprintf(os.Stderr, "Skipping %d: last got %d\n", i, i-failed)
			i++
			continue
		}

		if res.StatusCode == http.StatusOK {

			var dataJ dataReceived
			if err := json.NewDecoder(res.Body).Decode(&dataJ); err != nil {
				fmt.Fprintf(os.Stderr, "Error decoding data comic num %d\n", i)
				i++
				continue
			}

			comicJ := comicInfo{
				Link:       baseUrl + fmt.Sprint(dataJ.Num) + "/",
				Date:       dataJ.Month + "/" + dataJ.Day + "/" + dataJ.Year,
				Title:      dataJ.Title,
				Transcript: dataJ.Transcript,
			}
			comics = append(comics, comicJ)

			fmt.Printf("processing: got %d\n", i)
			failed = 0
		}
		res.Body.Close()
		i++
	}
	fmt.Printf("read %d comics", i-failed)

	reader, err := os.Create(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening file")
		os.Exit(-1)
	}

	defer reader.Close()

	if err := json.NewEncoder(reader).Encode(comics); err != nil {
		fmt.Fprintln(os.Stderr, "Error encoding data")
		os.Exit(-1)
	}

}
