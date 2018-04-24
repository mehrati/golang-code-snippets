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
	std1 := Student{"Ali", "Salimy", 12}
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err.Error())
	}
	encoder := gob.NewEncoder(conn)
	encoder.Encode(std1)
	conn.Close()

}
