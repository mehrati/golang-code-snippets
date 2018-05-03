package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i <= 100; i++ {
		ProgessBar(i)
		time.Sleep(time.Second / 10)
	}
}
func ProgessBar(percent int) {
	hash := "####################"
	space := "                    "
	lspace := len(space)
	max := percent / 5
	for i := 0; i <= max; i++ {
		fmt.Printf("\r %d%% [%s%s] ", percent, hash[:i], space[:lspace-i])
	}
}
