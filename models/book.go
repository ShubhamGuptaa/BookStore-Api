package models

import 
(
	"gopkg.in/mgo.v2/bson"
)
type Book struct {
	Id bson.ObjectId     `json: "id" bson:"_id"`
	Name string 		 `json:"name" bson:"name"`
	Ratings string		 `json:"ratings" bson:"ratings"`
	Description string   `json:"description" bson:"description"`
}
