package booking

import (
	"fmt"

	"github.com/Silvin/booking/io"
	"github.com/Silvin/booking/models"

	"gopkg.in/mgo.v2/bson"
)

type Booking struct {
	db io.MgoConn
}

//Init the connection with the Database
func (this *Booking) Init(DBHost string, DBName string) error {
	err, message := this.db.New(DBHost, DBName)
	if err != nil {

		fmt.Println(message)

	}
	defer this.db.Close()
	return err

}

//Add a new PaymentType
func (this *Booking) AddPaymentType(name string) (models.PaymentType, error) {
	payment := models.PaymentType{bson.NewObjectId(), name}
	err := this.db.DB.C(payment.GetTableName()).Insert(payment)
	return payment, err

}

//Add a new Client
func (this *Booking) AddClient(name string, username string, password string, comission models.Comission) (models.Client, error) {

	client := models.Client{bson.NewObjectId(), name, username, password, comission.Id}
	err := this.db.DB.C(client.GetTableName()).Insert(client)
	return client, err

}

//Add a new Event
func (this *Booking) AddEvent(name string, clientId bson.ObjectId, comission models.Comission) (models.Event, error) {

	event := models.Event{bson.NewObjectId(), name, clientId, comission}
	err := this.db.DB.C(event.GetTableName()).Insert(event)
	return event, err
}

//Create a new Comission Object :)
func (this *Booking) AddCommission(name string, value float32, isPercent bool, isDefault bool) (models.Comission, error) {
	comission := models.Comission{bson.NewObjectId(), name, value, isPercent, isDefault}
	err := this.db.DB.C(comission.GetTableName()).Insert(comission)
	return comission, err
}

//Remove a payment
func (this *Booking) RemovePayment(payment models.PaymentType) error {
	return this.db.Delete(payment.GetTableName(), bson.M{"_id": payment.Id.Hex()})
}

//Remove a Client
func (this *Booking) RemoveClient(client models.Client) error {
	return this.db.Delete(client.GetTableName(), bson.M{"_id": client.Id.Hex()})
}

//Remove a Event
func (this *Booking) RemoveEvent(event models.Event) error {
	return this.db.Delete(event.GetTableName(), bson.M{"_id": event.Id.Hex()})
}

//Update a payment Data
func (this *Booking) UpdatePayment(payment models.PaymentType) error {

	return this.db.Update(
		payment.GetTableName(),
		bson.M{"_id": payment.Id.Hex()},
		bson.M{"$set": bson.M{"name": payment.Name}})
}

//Update client data
func (this *Booking) UpdateClient(client models.Client) error {
	return this.db.Update(
		client.GetTableName(),
		bson.M{"_id": client.Id.Hex()},
		bson.M{"$set": bson.M{
			"name":      client.Name,
			"username":  client.Username,
			"password":  client.Password,
			"comission": client.Comission}})
}

//update event data
func (this *Booking) UpdateEvent(event models.Event) error {

	return this.db.Update(
		event.GetTableName(),
		bson.M{"_id": event.Id.Hex()},
		bson.M{"$set": bson.M{
			"name":      event.Name,
			"client_id": event.Client_id,
			"comission": event.Comission}})
}

//Get All payments type
func (this *Booking) GetAllPaymentsType() ([]models.PaymentType, error) {

	payments := []models.PaymentType{}
	payment := models.PaymentType{}
	err := this.db.GetAll(payment.GetTableName(), &payments)
	return payments, err

}

//Get All clients
func (this *Booking) GetAllClients() ([]models.Client, error) {

	clients := []models.Client{}
	client := models.Client{}
	err := this.db.GetAll(client.GetTableName(), &clients)
	return clients, err

}

//Get All Events
func (this *Booking) GetAllEvents() ([]models.Event, error) {

	events := []models.Event{}
	event := models.Event{}
	err := this.db.GetAll(event.GetTableName(), &events)
	return events, err
}

//Get All Comissions
func (this *Booking) GetAllComission() ([]models.Comission, error) {

	comissions := []models.Comission{}
	comission := models.Comission{}
	err := this.db.GetAll(comission.GetTableName(), &comissions)
	return comissions, err
}

//Get Default Commision
func (this *Booking) GetDefaultComission() (models.Comission, error) {

	comissions := []models.Comission{}
	comission := models.Comission{}
	err := this.db.Find(comission.GetTableName(), bson.M{"is_default": true}, &comissions)

	if len(comissions) > 0 {
		comission = comissions[0]
	} else {

		err = this.db.GetAll(comission.GetTableName(), &comissions)
		if err == nil {

			if len(comissions) > 0 {
				comission = comissions[0]
			}

		}
	}

	return comission, err
}

//Get All Events by clients
func (this *Booking) GetEventsByClient(client models.Client) ([]models.Event, error) {

	events := []models.Event{}
	event := models.Event{}
	err := this.db.Find(event.GetTableName(), bson.M{"clien_id": client.Id}, &events)
	return events, err
}

//Get real cost from a event
func (this *Booking) CalculateCost(
	event models.Event,
	payment models.PaymentType,
	quantity int,
	ticketCost float32) (models.BookingCost, error) {

	return event.GetTotal(payment.Id, quantity, ticketCost)
}

//Remove the database
func (this *Booking) RemoveDB() {

	this.db.DB.DropDatabase()

}
