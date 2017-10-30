package models

import "gopkg.in/mgo.v2/bson"

type Client struct {
	Id            bson.ObjectId `bson:"_id,omitempty"       json:"Id"`
	Name          string        `bson:"name"            	json:"Name"`
	Username      string        `bson:"username"        	json:"Username"`
	BaseComission []Comission   `bson:"base_comission"  	json:"BaseCommision"`
}

func (this *Client) GetTableName() string {

	return "client"
}
