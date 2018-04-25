package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, e := os.Open("main.go")
	if e != nil {

	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
