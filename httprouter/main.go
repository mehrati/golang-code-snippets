package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/hello/:name", hello)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Index")
}
func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello %s", ps.ByName("name"))
}
