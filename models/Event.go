package models

import "gopkg.in/mgo.v2/bson"

type Event struct {
	Id        bson.ObjectId `bson:"_id,omitempty"       json:"Id"`
	Name      string        `bson:"name"            	json:"Name"`
	Client    Client        `bson:"client"        		json:"Username"`
	Comission []Comission   `bson:"comission"  			json:"BaseCommision"`
}

func (this *Event) GetTableName() string {

	return "event"
}
