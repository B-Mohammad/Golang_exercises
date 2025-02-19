package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/"
	res, err := http.Get(url + "/todos/" + os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(-1)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(-1)
	}

	var item todo

	err = json.Unmarshal(data, &item)
	if err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		os.Exit(-1)
	}

	fmt.Println(item)
}
