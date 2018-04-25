package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("loging.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	log.Println("loging")
}
