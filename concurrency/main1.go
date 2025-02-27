package main

// import (
// 	"log"
// 	"net/http"
// 	"time"
// )

// type respond struct {
// 	url     string
// 	latency time.Duration
// 	err     error
// }

// func getData(url string, result chan<- respond) {
// 	start := time.Now()

// 	if res, err := http.Get(url); err != nil {
// 		result <- respond{url: url, latency: 0, err: err}
// 	} else {
// 		t := time.Since(start).Round(time.Millisecond)
// 		result <- respond{url: url, latency: t, err: nil}
// 		res.Body.Close()
// 	}

// }

// func main() {
// 	s := time.Now()

// 	urls := []string{
// 		"https://www.golang.org/",
// 		"https://www.google.com/",
// 		"https://www.amazon.com/",
// 		"https://ui.ac.ir",
// 		"https://tamin.ir",
// 		"https://w3schools.com"}

// 	ch := make(chan respond)

// 	for _, u := range urls {
// 		go getData(u, ch)
// 	}

// 	for range urls {
// 		res := <-ch
// 		if res.err != nil {
// 			log.Printf("%-50s %s\n", res.url, res.err)
// 		} else {
// 			log.Printf("%-50s %s\n", res.url, res.latency)

// 		}
// 	}

// 	log.Println(time.Since(s).Round(time.Millisecond))

// }
