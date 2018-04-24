package main

import (
	"fmt"
	"net/url"
)

func main() {
	var ur *url.URL
	ur, err := url.Parse("www.google.com/")
	if err != nil {
		fmt.Println(err.Error())
	}
	parameters := url.Values{}

	parameters.Add("k", "key")
	parameters.Add("n", "name")

	ur.RawQuery = parameters.Encode()
	fmt.Println(ur)
}
