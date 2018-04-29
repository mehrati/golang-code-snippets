package main

import (
	"fmt"
	"regexp"
)

func main() {
	isLetter := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	fmt.Println(isLetter("GHF^"))

}
