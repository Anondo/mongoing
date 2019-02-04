package main

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func findAll(clctn string) ([]Message, error) {
	var ms []Message
	collection := getCollection(clctn)

	cur, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var m Message
		err := cur.Decode(&m)
		if err != nil {
			return nil, err
		}

		ms = append(ms, m)
	}

	return ms, nil
}

func findOne(clctn string, fltr bson.M) (*Message, error) {
	collection := getCollection(clctn)

	var m *Message
	err := collection.FindOne(ctx, fltr).Decode(&m)

	if err != nil {
		return nil, err
	}

	return m, nil
}

func insert(clctn string, v bson.M) error {
	collection := getCollection(clctn)

	res, err := collection.InsertOne(ctx, v)

	if err != nil {
		return err
	}

	fmt.Println("One New Row Inserted Successfully, ID: ", res.InsertedID)
	return nil

}

func updateOne(clctn string, fltr bson.M, qry bson.M) error {
	collection := getCollection(clctn)

	res, err := collection.UpdateOne(ctx, fltr, qry)

	if err != nil {
		return err
	}

	fmt.Println("Modified documents: ", res.ModifiedCount)
	return nil
}

func deleteOne(clctn string, fltr bson.M) error {
	collection := getCollection(clctn)

	res, err := collection.DeleteOne(ctx, fltr)

	if err != nil {
		return err
	}

	fmt.Println("Delete messages: ", res.DeletedCount)
	return nil
}
