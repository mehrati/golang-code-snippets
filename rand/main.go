package main

import (
	"time"
	"fmt"
	"math/rand"
)
const rg = 100
func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(rg))
}
