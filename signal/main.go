package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var done chan struct{}

func main() {
	done = make(chan struct{})
	go FreeResource()
	<-done
	fmt.Println("done")
}
func FreeResource() {

	sig := make(chan os.Signal)
	signal.Notify(sig,
		syscall.SIGINT,
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	for {
		select {
		case <-sig:
			fmt.Println("free rs")
			close(done)
		}
	}
}
