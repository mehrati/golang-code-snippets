package main

import (
	"os/exec"
	"os"
)

func main() {
	out,err := exec.Command("clear").Output()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	os.Stdout.Write(out)
}
