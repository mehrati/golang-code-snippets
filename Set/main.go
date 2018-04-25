package main

import (
	"strings"
	"fmt"

	"gopkg.in/fatih/set.v0"
)

func main() {
	s := set.New()
	s.Add("fmt")
	s.Add("log")
	s.Add("log")
	s.Add("log")
	s.Add("fmt")
	s.Add("fmt ")
	l := s.List()
	for _, i := range l {
		fmt.Println(i.(string))
	}

}

func TrimImport(str string)string{
	s := ""
	s = strings.TrimLeft(str, s[:strings.Index(s, "\"")])
	s = strings.TrimRight(s, s[strings.LastIndex(s,"\""):])
	return s
}