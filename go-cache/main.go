package main

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"
)

func main() {
	fmt.Println("")
	c := cache.New(time.Minute, time.Minute*2)
	c.Set("key1","value1",cache.DefaultExpiration)
	c.Set("key2","value2",cache.NoExpiration)
	time.Sleep(time.Minute * 2)
	fmt.Println(c.Get("key1"))
	fmt.Println(c.Get("key2"))
}
