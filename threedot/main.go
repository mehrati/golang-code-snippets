package main

import (
	"fmt"
)

func main() {
	arr := [6]string{"ali","hassan","mohsen","sasan","saeed","maryam"}
	printNames(arr[:]...)
}

func printNames(names ...string){
	for i:=range names{
		fmt.Println(names[i])
	}
}