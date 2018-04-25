package main

import (
	"fmt"
)

func main() {
	var a int = 10
	var b int = 20
	fmt.Printf("befor swap a => %d , b => %d \n", a, b)
	Swap(&a, &b)
	fmt.Printf("after swap a => %d , b => %d \n", a, b)

}

func Swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
