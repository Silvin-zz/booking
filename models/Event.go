package models

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

type Event struct {
	Id        bson.ObjectId `bson:"_id,omitempty"       json:"Id"`
	Name      string        `bson:"name"                json:"Name"`
	Client_id bson.ObjectId `bson:"client_id"           json:"ClientId"`
	Comission []Comission   `bson:"comission"           json:"Commision"`
}

type BookingCost struct {
	Quantity                 int
	Subtotal                 float32
	TotalBase                float32
	TotalPercentageCommision float32
	Cost                     float32
}

func (this *Event) GetTableName() string {

	return "event"
}

//GET THE TOTAL COST FROM A TICKET.
//The function does not know the ticket cost we assumed that the cost came from another design
//This because we can have an event with many types of costs, zones and stocks.
func (this *Event) GetTotal(PaymentType bson.ObjectId, quantity int, ticketCost float32) (BookingCost, error) {

	//Find the commision
	eventComission, err := this.getPaymentCommision(PaymentType)

	if err == nil {

		// Calculate the cost from total tickets
		subtotal := float32(quantity) * ticketCost
		comissionBase := float32(quantity) * eventComission.Base
		comissionPercentage := (subtotal + comissionBase) * (eventComission.Percentage / 100)
		cost := subtotal + comissionBase + comissionPercentage

		total := BookingCost{quantity, subtotal, comissionBase, comissionPercentage, cost}

		return total, nil

	}
	return BookingCost{}, errors.New("Error getting the commision data")

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
