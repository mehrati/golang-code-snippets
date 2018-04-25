package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	getIP()
}

func getIP() {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Fatal(err)
		}
		for _, addr := range addrs {
			fmt.Println(addr.String())
		}
	}
}
