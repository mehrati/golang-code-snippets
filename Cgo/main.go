package main

// #include<float.h>
import "C"
import "fmt"

func main() {
	fmt.Println(C.FLT_MAX)
}
