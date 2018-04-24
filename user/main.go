package main

import (
  "fmt"
  "os/user"
)

func main() {

  user, err := user.Current()
  if err != nil {
    fmt.Println(err.Error())
  }
  fmt.Printf("Hello %s \n",user.Username)
  fmt.Println(user.Uid)
}