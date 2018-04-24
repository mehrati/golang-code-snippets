package main

import (
	"fmt"
	"github.com/robfig/cron"
)

var done chan bool

func main() {
	done = make(chan bool,1)
	cr := cron.New()
	cr.AddFunc("@every 2s",SayHello)
	cr.AddFunc("59 * * * * *",SayBye)
	cr.Start()
	<-done
}

func SayHello(){
	fmt.Println("Hello !")
}
func SayBye(){
	fmt.Println("Bye !")
	done<-true
}