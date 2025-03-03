package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.Second * 2

	ch := time.NewTicker(ticker).C

	done := time.After(time.Second * 12)
	var i int
loop:
	for {
		select {
		case <-done:
			fmt.Println(i)
			i += 2
			break loop
		case <-ch:
			fmt.Println(i)
			i += 2
		}
	}
}
