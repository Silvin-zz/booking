package models

import "gopkg.in/mgo.v2/bson"

type Comission struct {
	Id          bson.ObjectId `bson:"_id,omitempty"    		json:"Id"`
	PaymentType PaymentType   `bson:"payment_type"    		json:"Type"`
	Base        float32       `bson:"base"	    			json:"Base"`
	Percentage  float32       `bson:"percentage"			json:"Percentage"`
}

func (this *Comission) GetTableName() string {

	return "comission"
}
