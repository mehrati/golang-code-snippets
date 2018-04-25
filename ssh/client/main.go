package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

var (
	Ip       = ""
	Port     = "22"
	password = ""
	Command  = "ls -la"
)

func main() {
	sshConf := &ssh.ClientConfig{
		User: "user",
		Auth: []ssh.AuthMethod{
			ssh.Password(password)},
	}
	sshConf.HostKeyCallback = ssh.InsecureIgnoreHostKey()
	cli, err := ssh.Dial("tcp", Ip+":"+Port, sshConf)
	if err != nil {
		log.Fatal(err)
	}
	s, err := cli.NewSession()

	b, err := s.CombinedOutput(Command)
	fmt.Println(string(b))
	cli.Close()
}
