package main

import (
	"fmt"

	"github.com/atotto/clipboard"
)

func main() {
	str, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
	clipboard.WriteAll("ok i read !!!")
}
