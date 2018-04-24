package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("loging.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	log.SetOutput(file)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	log.Println("loging")
}
