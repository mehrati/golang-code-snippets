package main

import (
	"fmt"
)

func main() {
	ch := OutChan()
	for c := range ch {
		fmt.Println(c)
	}
}

func OutChan() chan int {
	ch := make(chan int, 1)
	i := 0
	go func() {
		for {
			i++
			ch <- i
		}
	}()
	return ch
}
