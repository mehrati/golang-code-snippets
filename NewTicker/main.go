package main

import (
	"fmt"
	"time"
)

func main() {
	tiker := time.NewTicker(time.Second)
	for c := range tiker.C {
		fmt.Println(c.Second())
	}
}
