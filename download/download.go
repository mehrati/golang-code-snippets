package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

func main() {
	link := ""
	name := path.Base(link)
	data, err := Download(link)
	if err != nil {
		log.Println(err.Error())
	}
	err = ioutil.WriteFile(name, data, 0755)
	if err != nil {
		fmt.Print(err.Error())
	}

}

func Download(url string) (data []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)
	// n, err := io.Copy(output, response.Body)
	if err != nil {
		log.Printf(err.Error())
		return nil, err
	}
	return data, nil

}
