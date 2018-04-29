package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", Hello)
	router.Methods("GET").Path("/how").Name("how").HandlerFunc(How)
	router.HandleFunc("/bye", Bye)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
	fmt.Fprintln(w, "Hello")
}
func How(w http.ResponseWriter, r *http.Request) {
	fmt.Println("How Are You ?")
	fmt.Fprintln(w, "How Are You ?")
}
func Bye(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Bye")
	fmt.Fprintln(w, "Bye")
}
