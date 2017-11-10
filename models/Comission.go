package models

import "gopkg.in/mgo.v2/bson"

type Comission struct {
	Id        bson.ObjectId `bson:"_id,omitempty"    json:"Id"`
	Name      string        `bson:"name"             json:"name"`
	Value     float32       `bson:"value"            json:"Value"`
	IsPercent bool          `bson:"is_percent"       json:"IsPercent"`
	IsDefault bool          `bson:"is_default"       json:"IsDefault"`
}

func (this *Comission) GetTableName() string {

	return "comission"
}
