package models

import "gopkg.in/mgo.v2/bson"

type Event struct {
	Id        bson.ObjectId `bson:"_id,omitempty"   json:"Id"`
	Name      string        `bson:"name"            json:"Name"`
	Client_id bson.ObjectId `bson:"client_id"       json:"ClientId"`
	Comission Comission     `bson:"comission"       json:"Commision"`
}

type BookingCost struct {
	Quantity       int
	Subtotal       float32
	TotalComission float32
	Cost           float32
}

func (this *Event) GetTableName() string {

	return "event"
}

//GET THE TOTAL COST FROM A TICKET.
//The function does not know the ticket cost we assumed that the cost came from another design
//This because we can have an event with many types of costs, zones and stocks.
func (this *Event) GetTotal(PaymentType bson.ObjectId, quantity int, ticketCost float32) (BookingCost, error) {

	//Find the commision
	eventComission := this.Comission
	extraCost := eventComission.Value

	if eventComission.IsPercent == true {
		extraCost = (eventComission.Value / 100) * ticketCost
	}

	// Calculate the cost from total tickets
	subtotal := float32(quantity) * ticketCost
	comissionTotal := float32(quantity) * extraCost
	cost := subtotal + comissionTotal

	total := BookingCost{quantity, subtotal, comissionTotal, cost}

	return total, nil

}
