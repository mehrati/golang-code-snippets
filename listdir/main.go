package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	listdirs := []string{wd}
	i := 0
	for {
		dirs, err := GetListDir(listdirs[i])
		if err != nil {
			break
		}
		for _, dir := range dirs {
			listdirs = append(listdirs, dir)
		}
		i++
		if i == len(listdirs) {
			break
		}
	}
	ShowDirs(listdirs)
}

func GetListDir(dir string) ([]string, error) {
	listdir := []string{}
	list, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, l := range list {
		if l.IsDir() {
			listdir = append(listdir, dir+"/"+l.Name())
		}
	}
	return listdir, nil
}

func ShowDirs(dirs []string) {
	for _, i := range dirs {
		fmt.Println(i)
	}
}
