package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	session, err := mgo.Dial(os.Getenv("MONGO_PORT_27017_TCP_ADDR"))
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")
	err = c.Insert(&Person{"Ale", "481823"},
		&Person{"Cla", "3421323"})
	if err != nil {
		log.Fatal(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Phone:", result.Phone)
}
