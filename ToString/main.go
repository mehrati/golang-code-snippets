package main

import "fmt"

type Message string

func (msg Message) String() string { return string("**( " + msg + " )**") }
func main() {
	msg := Message("Golang is good language")
	fmt.Println(msg)
}
