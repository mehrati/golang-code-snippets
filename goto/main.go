package main

import "fmt"

func main() {
	i := 0
START:
	fmt.Println(i)
	i++
	if i > 100 {
		goto END
	}
	goto START
	fmt.Println("never printed")
END:
	fmt.Println("End loop")
}
