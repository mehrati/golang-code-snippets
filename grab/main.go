package main

import (
	"fmt"
	"time"

	"github.com/cavaliercoder/grab"
)

func main() {
	client := grab.NewClient()
	requst, err := grab.NewRequest(".", "https://www.tutorialspoint.com/go/go_tutorial.pdf")
	if err != nil {
		fmt.Println(err)
	}
	response := client.Do(requst)
	fmt.Printf("size = %.2f Mb \n", float64(response.Size)/1000000)
	tiker := time.NewTicker(50000)
loop:
	for {
		select {
		case <-tiker.C:
			fmt.Println(response.BytesComplete())
			ProgessBar(int(100 * response.Progress()))
		case <-response.Done:
			fmt.Println()
			fmt.Printf("done. \n")
			break loop
		}
	}
}
func ProgessBar(percent int) {
	hash := "####################"
	space := "                    "
	lspace := len(space)
	max := percent / 5
	for i := 0; i <= max; i++ {
		fmt.Printf("\r %d%% [%s%s] ", percent, hash[:i], space[:lspace-i])
	}
}
