package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("")
	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	clinet := &http.Client{}
	resp, err := clinet.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(b))
}
