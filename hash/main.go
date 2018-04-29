package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := "secretpass"
	hash := CreateHashPassword(pass)
	fmt.Println(hash)
	re := ComparePassword(hash, pass)
	fmt.Println(re)

}
func CreateHashPassword(input string) string {

	out, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.MinCost)
	if err != nil {
		log.Fatalln(err)
	}
	return string(out)
}
func ComparePassword(hash, pass string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
