package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

var signalChan chan os.Signal
var cleanDone chan bool

func main() {
	fmt.Println("Start Prosess ", os.Getpid())
	signalChan = make(chan os.Signal)
	cleanDone = make(chan bool, 1)
	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		for sig := range signalChan {
			fmt.Printf("Signal %s Received \n", sig)
			Clean()
			cleanDone <- true
		}
	}()
	<-cleanDone
}

func Clean() {
	b, e := exec.Command("clear").Output()
	if e != nil {
		fmt.Println(e.Error())
	}
	os.Stdout.Write(b)
}
