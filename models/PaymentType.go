package models

import "gopkg.in/mgo.v2/bson"

type PaymentType struct {
	Id   bson.ObjectId `bson:"_id,omitempty"   	json:"Id"`
	Name string        `bson:"name"    			json:"Id"`
}

func (this *PaymentType) GetTableName() string {

	return "paymentType"
}
