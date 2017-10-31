package main

import (
	"comission/io"
	"comission/models"
	"fmt"

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
func (this *Booking) AddClient(name string, username string, password string, comissions []models.Comission) (models.Client, error) {

	client := models.Client{bson.NewObjectId(), name, username, password, comissions}
	err := this.db.DB.C(client.GetTableName()).Insert(client)
	return client, err

}

//Add a new Event
func (this *Booking) AddEvent(name string, clientId bson.ObjectId, comissions []models.Comission) (models.Event, error) {

	event := models.Event{bson.NewObjectId(), name, clientId, comissions}
	err := this.db.DB.C(event.GetTableName()).Insert(event)
	return event, err
}

//Create a new Comission Object :)
func (this *Booking) AddCommission(paymentType models.PaymentType, base float32, percentage float32) models.Comission {
	comission := models.Comission{bson.NewObjectId(), paymentType, base, percentage}
	return comission
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
			"name":           client.Name,
			"username":       client.Username,
			"password":       client.Password,
			"base_comission": client.BaseComission}})
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

//Get All Events by clients
func (this *Booking) GetEventsByClient(client models.Client) ([]models.Event, error) {

	events := []models.Event{}
	event := models.Event{}
	err := this.db.Find(event.GetTableName(), bson.M{"clien_id": client.Id}, &events)
	return events, err
}

//Get real cost from a event
func (this *Booking) CalculateCost(event models.Event, payment models.PaymentType, quantity int, ticketCost float32) (models.BookingCost, error) {

	return event.GetTotal(payment.Id, quantity, ticketCost)
}

//Remove the database
func (this *Booking) RemoveDB() {

	this.db.DB.DropDatabase()

}

func main() {

	//Set the database connection
	booking := Booking{}
	booking.Init("127.0.0.1:27017", "test") //After we going to implement auth access
	booking.RemoveDB()                      //Remove the database if exists

	//Define the payments type
	paypal, _ := booking.AddPaymentType("PayPal")
	mastercard, _ := booking.AddPaymentType("Mastercard")

	//Generate a Base commision
	comissions := []models.Comission{}
	comissions = append(comissions, booking.AddCommission(paypal, 5, 1))
	comissions = append(comissions, booking.AddCommission(mastercard, 4, 1.5))

	//Create a client with base comission
	client, _ := booking.AddClient("Client 1", "client1", "pass", comissions)

	//Modify the Commision Base for a client event

	//Modify  comission for paypal
	comissions[0].Base = 4
	comissions[0].Percentage = 1

	//Modify comission for mastercard
	comissions[1].Base = 3
	comissions[1].Percentage = 0.75

	//Add the new event
	event, _ := booking.AddEvent("Corona Capital", client.Id, comissions)

	//Get the a ticket cost from paypal
	paypalCost, _ := booking.CalculateCost(event, paypal, 1, 100)

	//Print the result with paypal
	fmt.Println("_______________________ PAYPAL _____________________")
	fmt.Println(fmt.Sprintf("Quantity: %d", paypalCost.Quantity))
	fmt.Println(fmt.Sprintf("Subtotal: %.2f", paypalCost.Subtotal))
	fmt.Println(fmt.Sprintf("Base Comission: %.2f", paypalCost.TotalBase))
	fmt.Println(fmt.Sprintf("Percentage Comission: %.2f", paypalCost.TotalPercentageCommision))
	fmt.Println(fmt.Sprintf("Total: %.2f", paypalCost.Cost))

	//Get the a ticket cost from paypal
	masterCost, _ := booking.CalculateCost(event, mastercard, 1, 100)

	//Print the result with mastercard
	fmt.Println("_______________________ MASTERCARD _____________________")
	fmt.Println(fmt.Sprintf("Quantity: %d", masterCost.Quantity))
	fmt.Println(fmt.Sprintf("Subtotal: %.2f", masterCost.Subtotal))
	fmt.Println(fmt.Sprintf("Base Comission: %.2f", masterCost.TotalBase))
	fmt.Println(fmt.Sprintf("Percentage Comission: %.2f", masterCost.TotalPercentageCommision))
	fmt.Println(fmt.Sprintf("Total: %.2f", masterCost.Cost))

}
