package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type data struct {
	name, price string
}

var testData = []data{
	{"apple", "1.1"},
	{"shirt", "2.2"},
	{"orange", "3.3"},
	{"shoes", "4.4"},
	{"socks", "200.22"},
}

func reqSender(endpoint, params string) {

	url := "http://172.31.141.0:8080/" + endpoint + "?" + params
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	fmt.Fprintf(os.Stdout, "%s with params %s (status: %d) ---> NO Error\n", endpoint, params, res.StatusCode)
}

func addReq() {
	for {
		for _, item := range testData {
			reqSender("add", "item="+item.name+"&price="+item.price)
		}
	}
}

func updateReq() {
	for {
		for _, item := range testData {
			reqSender("update", "item="+item.name+"&price="+item.price)
		}
	}
}

func deleteReq() {
	for {
		for _, item := range testData {
			reqSender("delete", "item="+item.name)
		}
	}
}

func main() {

	go addReq()
	go updateReq()
	go deleteReq()

	time.Sleep(10 * time.Second)

}
