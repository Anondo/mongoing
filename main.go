package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"
)

func init() {
	loadConfig()
	connectMongo()
}

func main() {

	fmt.Println("Getting all message...")
	ms, err := findAll("message")

	if err != nil {
		log.Fatal(err)
	}

	showFoundData(ms...)

	fmt.Println("Getting one message...")
	m, _ := findOne("message", bson.M{"by": "Anondo"})

	showFoundData(*m)

	fmt.Println("Inserting data...")

	err = insert("message", bson.M{"msg": "New Message", "by": "John Doe"})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Updating one message...")
	filter := bson.M{"by": "John Doe"}
	qry := bson.M{
		"$set": bson.M{
			"msg": "Latest Message",
		},
	}
	err = updateOne("message", filter, qry)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleting one message...")
	err = deleteOne("message", bson.M{"by": "John Doe"})

	if err != nil {
		log.Fatal(err)
	}

}

func showFoundData(ms ...Message) {
	for i, m := range ms {
		fmt.Println("Message No:", i+1)
		fmt.Println("---------------------")
		fmt.Println("Message: ", m.Msg)
		fmt.Println("By: ", m.By)
		fmt.Println("---------------------")

	}
}
