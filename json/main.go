package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Student struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Id       int    `json:"id"`
	Age      int    `json:"age"`
}
type Students struct {
	Students []Student `json:"students"`
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
func CreateFile(name string) io.WriteCloser {
	file, err := os.Create(name)
	CheckErr(err)
	return file
}
func main() {
	st := Students{[]Student{
		Student{"Maryam", "hoseiny", 1, 19},
		Student{"Hassan", "jalily", 2, 23},
		Student{"Sasan", "moradi", 3, 17},
	}}
	fw := CreateFile("strudents.json")
	defer fw.Close()
	encoder := json.NewEncoder(fw)
	CheckErr(encoder.Encode(st))
	fr, err := os.Open("strudents.json")
	CheckErr(err)
	defer fr.Close()
	decoder := json.NewDecoder(fr)
	s := Students{}
	decoder.Decode(&s)
	fmt.Println(s)

}
