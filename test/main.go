package main

import (
	"fmt"
)

func main() {

	str := RepeatChar(10, 'e')
	fmt.Println(str)
	
}

func RepeatChar(n int, r rune) string {
	s := ""
	for i := 0; i < n; i++ {
		s += string(r)
	}
	return s
}
