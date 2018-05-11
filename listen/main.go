package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", Handler)
	err = http.Serve(listener, mux)
	if err != nil {
		fmt.Println(err)
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}
