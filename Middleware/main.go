package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.AddCookie(Cookie())
		next(w, r)
	})
}

func Hello(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("Key")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(c.Value)
	}

	fmt.Fprintln(w, "hello")
}

func main() {

	http.HandleFunc("/",Middleware(Hello))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Cookie() *http.Cookie {
	exp := time.Now().Add(30 * 24 * time.Hour)
	cookie := http.Cookie{Name: "Key", Value: "Value", Expires: exp}
	return &cookie
}
