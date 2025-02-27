package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// type userId chan int

// func (id userId) handler(w http.ResponseWriter, r *http.Request) {
// 	u := <-(id)
// 	fmt.Fprintf(w, "%d", u)
// }

// func counter(ch chan<- int) {
// 	for i := 1; ; i++ {
// 		ch <- i
// 	}
// }

// func main() {
// 	var ch userId = make(chan int)
// 	go counter(ch)

// 	http.HandleFunc("/", ch.handler)

// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
