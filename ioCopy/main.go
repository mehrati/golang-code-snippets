package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	data := strings.NewReader("data plan text string ")
	file, err := os.Create("data.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	io.Copy(file, data)
}
