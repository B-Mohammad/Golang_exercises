package main

// import (
// 	"log"
// 	"time"
// )

// func main() {
// 	chans := []chan int{
// 		make(chan int),
// 		make(chan int),
// 	}

// 	for i, v := range chans {
// 		go func(ch chan int, i int) {
// 			for {
// 				time.Sleep(time.Duration(i) * time.Second)
// 				ch <- i
// 			}
// 		}(v, i+1)
// 	}

// 	for i := 0; i < 12; i++ {
// 		select {
// 		case r0 := <-chans[0]:
// 			log.Printf("goroutine %d %d", i, r0)

// 		case r1 := <-chans[1]:
// 			log.Printf("goroutine %d %d", i, r1)
// 		}
// 	}
// }
