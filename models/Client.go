package models

import "gopkg.in/mgo.v2/bson"

type Client struct {
	Id        bson.ObjectId `bson:"_id,omitempty"       json:"Id"`
	Name      string        `bson:"name"            	json:"Name"`
	Username  string        `bson:"username"        	json:"Username"`
	Password  string        `bson:"password"        	json:"-"`
	Comission bson.ObjectId `bson:"comission"  			json:"Commision"`
}

func (this *Client) GetTableName() string {

	return "client"
}
