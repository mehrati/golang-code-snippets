package main

import "fmt"

func main() {
  slice := []int{1, 2, 3, 4}
  ChangeSlice(slice)
  fmt.Println(slice)
}

func ChangeSlice(s []int) {
  s[0] = 0
  s = append(s, 5)
}