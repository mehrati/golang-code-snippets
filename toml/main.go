package main

import (
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

type Person struct {
	Name   string `toml:"name"`
	Famliy string `toml:"family"`
}
type Persons struct {
	Persons []Person `toml:"person"`
}

func main() {
	var p Persons
	by, err := ioutil.ReadFile("conf.toml")
	if err != nil {
		fmt.Println(err.Error())
	}

	metadata, err := toml.Decode(string(by), &p)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(p)
	fmt.Println(metadata.Keys())
}
