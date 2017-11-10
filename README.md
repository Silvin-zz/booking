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
    american, _ := booking.AddPaymentType("American Express")

    fmt.Println("... ... PAYMENTS ... ... ")
    fmt.Println(paypal.Name)
    fmt.Println(mastercard.Name)
    fmt.Println(american.Name)

    // //Generate a Base commision

    defaultComission, _ := booking.AddCommission("Default", 5, false, true) // Comission with value
    midComission, _ := booking.AddCommission("Medium", 2, true, true)       //Comission with percent
    premiumComission, _ := booking.AddCommission("Premium", 4, false, true) //Comission with percent

    fmt.Println("... ... CLIENTS ... ... ")
    fmt.Println(defaultComission.Name)
    fmt.Println(midComission.Name)
    fmt.Println(premiumComission.Name)

    // //Create a client with default comission
    defaultComission, _ = booking.GetDefaultComission()
    client, _ := booking.AddClient("Client 1", "client1", "pass", defaultComission)

    fmt.Println(client.Name)

    //Create a custom comission based on defaultComission
    customComission := defaultComission
    customComission.IsDefault = false
    customComission.Value = 2
    customComission.Name = "Custom lollapalooza"

    // Add the new event with default commission
    coronaCapital, _ := booking.AddEvent("Corona Capital", client.Id, midComission)
    viveLatino, _ := booking.AddEvent("Vive Latino", client.Id, premiumComission)
    lollapalooza, _ := booking.AddEvent("lollapalooza", client.Id, customComission)

    fmt.Println("... ... COMISSIONS ... ... ")
    fmt.Println(coronaCapital.Name)
    fmt.Println(viveLatino.Name)
    fmt.Println(lollapalooza.Name)

    //Calculate booking cost.

    result, _ := booking.CalculateCost(coronaCapital, mastercard, 1, 500)

    fmt.Println("... ... RESULT ... ... ")
    fmt.Println(fmt.Sprintf("Quantity: %d", result.Quantity))
    fmt.Println(fmt.Sprintf("Subtotal: %.2f", result.Subtotal))
    fmt.Println(fmt.Sprintf("Comission: %.2f", result.TotalComission))
    fmt.Println(fmt.Sprintf("Total: %.2f", result.Cost))

}



```

Regards

