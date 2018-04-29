package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hi", SayHi)
	mux.HandleFunc("/bye", SayBye)
	r := handlers.LoggingHandler(os.Stdout,mux)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func SayHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi")
}
func SayBye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bye")
}
