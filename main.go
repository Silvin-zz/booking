package main

import (
	"comission/io"
	"comission/models"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

func main() {

	db := io.MgoConn{}

	err, message := db.New("127.0.0.1:27017", "boletia")
	if err != nil {

		fmt.Println(message)
		return
	}

	//Init Data

	//Payment Types
	paypal := models.PaymentType{bson.NewObjectId(), "PayPal"}
	db.DB.C(paypal.GetTableName()).Insert(paypal)
	visa := models.PaymentType{bson.NewObjectId(), "Visa/MasterCard"}
	db.DB.C(visa.GetTableName()).Insert(visa)
	american := models.PaymentType{bson.NewObjectId(), "American Express"}
	db.DB.C(american.GetTableName()).Insert(american)

	//Default client comission
	clientComission := []models.Comission{}
	clientComission = append(clientComission, models.Comission{bson.NewObjectId(), paypal, 5, 1.5})
	clientComission = append(clientComission, models.Comission{bson.NewObjectId(), visa, 4, 2.0})
	clientComission = append(clientComission, models.Comission{bson.NewObjectId(), american, 6, 1.0})

	//Client
	client := models.Client{bson.NewObjectId(), "Client 1", "client1", clientComission}
	db.DB.C(client.GetTableName()).Insert(client)

	//Modify Client Comission
	clientComission[0].Base = 3
	clientComission[0].Percentage = 1.5
	clientComission[2].Base = 4
	clientComission[0].Percentage = 2.5
	//Event
	event := models.Event{bson.NewObjectId(), "Corona Capital", client, clientComission}
	db.DB.C(event.GetTableName()).Insert(event)

	fmt.Println("saludos")

}
