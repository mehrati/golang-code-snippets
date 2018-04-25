package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	done := make(chan bool,1)
	go Pro(ch)
	Con(ch, done)
	<-done
}

func Pro(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}
func Con(ch <-chan int, done chan bool) {
	for c := range ch {
		fmt.Println(c)
	}
	done <- true
}
