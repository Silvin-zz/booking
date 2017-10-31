# Booking
#### A simple comission library 



This library was written on.

  - Golang 
  - MongoDB
  

### Pre Installation

*You must have installed a mongodb server*

### Installation



```sh
$ go get github.com/Silvin/booking

```

How to Use ...

```go

package main

import (
	"fmt"
	"github.com/Silvin/booking"
	"github.com/Silvin/booking/models"
)

func main() {
	booking := booking.Booking{}
	booking.Init("127.0.0.1:27017", "test")
	booking.RemoveDB() //Remove the database if exists

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


```

Regards

