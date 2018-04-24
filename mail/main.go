package main

import (
	"fmt"
	"log"

	"gopkg.in/gomail.v2"
)

func main() {
	d := gomail.NewDialer("smtp.gmail.com", 587, "user@gmail.com", "password")
	m := gomail.NewMessage()
	m.SetHeader("From", "user@gmail.com")
	m.SetHeader("To", "address mail")
	m.SetHeader("Subject", "Hello")
	m.SetBody("text", "How Are You ???")
	err := d.DialAndSend(m)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Send mail...")

}
