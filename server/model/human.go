package model

import "github.com/globalsign/mgo/bson"

type Human struct {
	Fname string `json:"fname" bson:"fname" `
	Lname string `json:"lname" bson:"lname" `
	Age   string    `json:"age" bson:"age" `
	Id bson.ObjectId `json:"_id" bson:"_id"`
}


