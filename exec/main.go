package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	args := "-ltrh"
	path, err := exec.LookPath("ls")
	if err != nil {
		log.Fatal(err.Error())
	}
	cmd := exec.Command(path, args)
	res, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(res))
}
