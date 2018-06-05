package dao

import (
	"log"
  "../models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TopicsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "topics"
)

func (m *TopicsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *TopicsDAO) FindAll() ([]models.Topic, error) {
	var topics []models.Topic
	err := db.C(COLLECTION).Find(bson.M{}).All(&topics)
	return topics, err
}

func (m *TopicsDAO) FindById(id string) (models.Topic, error) {
	var topic models.Topic
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&topic)
	return topic, err
}

func (m *TopicsDAO) Insert(topic models.Topic) error {
	err := db.C(COLLECTION).Insert(&topic)
	return err
}

func (m *TopicsDAO) Delete(topic models.Topic) error {
	err := db.C(COLLECTION).Remove(&topic)
	return err
}

func (m *TopicsDAO) Update(topic models.Topic) error {
	err := db.C(COLLECTION).UpdateId(topic.ID, &topic)
	return err
}
