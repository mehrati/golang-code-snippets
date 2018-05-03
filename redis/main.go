package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	opt := redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	}
	client := redis.NewClient(&opt)
	defer client.Close()
	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Set("key11", 11, 0).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	res, err = client.Get("key11").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

}
