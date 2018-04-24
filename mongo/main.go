package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name     string `bson:"name"`
	LastName string `bson:"last_name"`
}

const (
	hosts      = "localhost:27017"
	database   = "db"
	username   = "user"
	password   = "pass"
	collection = "collection"
)

func main() {
	p1 := Person{"name", "lastname"}
	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer session.Close()
	coll := session.DB(database).C(collection)
	err = coll.Insert(p1)
	if err != nil {
		log.Fatal(err.Error())
	}
	p2 := Person{}
	err = coll.Find(bson.M{"name": "name"}).One(&p2)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(p2)
}
