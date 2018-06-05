package models

import "gopkg.in/mgo.v2/bson"


type Topic struct {
  ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	YearEstablished  int32        `bson:"yearEstablished" json:"yearEstablished"`
	Description string        `bson:"description" json:"description"`
}

type TopicInput struct {
  Name        string        `bson:"name" json:"name"`
	YearEstablished  string        `bson:"yearEstablished" json:"yearEstablished"`
	Description string        `bson:"description" json:"description"`
}
