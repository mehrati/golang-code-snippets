package main

import (
	"fmt"
)
func main() {
	fmt.Println("")
	slice := []int{0,2,3,4,5}
	ChangeSlice(slice)
	fmt.Println(slice)
}
func ChangeSlice(slc []int){
	slc[0]=1
	slc = append(slc,6)
}
