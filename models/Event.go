package models

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type A interface {
	GetAll(*mgo.Database) []A
}

type Event struct {
	Id        bson.ObjectId `bson:"_id,omitempty"       json:"Id"`
	Name      string        `bson:"name"                json:"Name"`
	Client_id bson.ObjectId `bson:"client_id"           json:"ClientId"`
	Comission []Comission   `bson:"comission"           json:"Commision"`
}

func (this *Event) GetTableName() string {

	return "event"
}

//GET THE TOTAL COST FROM A TICKET.
//The function does not know the ticket cost we assumed that the cost came from another design
//This because we can have an event with many types of costs, zones and stocks.
func (this *Event) GetTotal(PaymentType bson.ObjectId, quantity int, ticketCost float32) (float32, error) {

	//Find the commision
	eventComission, err := this.getPaymentCommision(PaymentType)

	if err == nil {

		// Calculate the cost from total tickets
		subtotal := float32(quantity) * ticketCost
		comissionBase := float32(quantity) * eventComission.Base
		comissionPercentage := subtotal * (eventComission.Percentage / 100)
		cost := subtotal + comissionBase + comissionPercentage

		return cost, nil

	}
	return 0, errors.New("Error getting the commision data")

}

//Get the paymentCommision into a Event Commision
func (this *Event) getPaymentCommision(id bson.ObjectId) (Comission, error) {

	for _, tmpComission := range this.Comission {

		if tmpComission.PaymentType.Id == id {
			return tmpComission, nil
		}
	}
	return Comission{}, errors.New("Invalid Payment Type")
}

// func (this *Event) GetAll(DB *mgo.Database) ([]A, error) {
// 	events := []Event{}

// 	DB.C(this.GetTableName()).Find(bson.M{}).All(&events)
// 	return events, nil
// }

// func (this *Event) Find(DB *mgo.Database, search bson.M) ([]Event, error) {

// }

// func (this *Event) Insert(DB *mgo.Database) error {

// }

// func (this *Event) Update(DB *mgo.Database, change bson.M) {

// }
