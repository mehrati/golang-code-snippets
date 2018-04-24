package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Student struct {
	Name     string
	LastName string
	Id       int
}

func main() {

	ls, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err.Error())
	}
	conn, err := ls.Accept()
	if err != nil {
		fmt.Println(err.Error())
	}
	std := handleConn(conn)
	fmt.Println(*std)
}

func handleConn(conn net.Conn) *Student {
	var std Student
	decoder := gob.NewDecoder(conn)
	decoder.Decode(&std)
	return &std
}
