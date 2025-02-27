package main

import (
	"fmt"
)

func generate(c1 chan int, limit int) {
	for i := 2; i <= limit; i++ {
		c1 <- i
	}
	close(c1)
}

func filter(source chan int, dst chan int, prime int) {
	for v := range source {
		if v%prime != 0 {
			dst <- v
		}
	}
	close(dst)
}

func sav(limit int) {
	ch := make(chan int)
	go generate(ch, limit)

	for {
		prime, ok := <-ch
		if !ok {
			break
		}

		ch2 := make(chan int)
		go filter(ch, ch2, prime)
		ch = ch2

		fmt.Print(prime, " ")

	}
	fmt.Println()
}

func main() {

	sav(10)

}
