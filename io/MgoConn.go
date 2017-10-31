package io

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var session *mgo.Session

type MgoConn struct {
	DB   *mgo.Database
	Conn *mgo.Session
}

//Create a MongoDB Connection
func (this *MgoConn) New(url string, database string) (error, string) {

	//we can use a auth parameters in production
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{url},
		Database: database,
		Timeout:  10 * time.Second,
	}

	sess, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		return err, fmt.Sprintf("The connect with mongoDB %s database in %s was faild", database, url)
	}

	session = sess
	session.SetMode(mgo.Monotonic, true)
	this.Conn = session.Copy()
	this.DB = this.Conn.DB(database)
	return nil, fmt.Sprintf("The connect with mongoDB %s database in %s was success", database, url)

}

//Close a MongoDB Connection
func (this *MgoConn) Close() {
	session.Close()
}

//Get All Elements in the collection
func (this *MgoConn) GetAll(collectionName string, result interface{}) error {
	return this.DB.C(collectionName).Find(bson.M{}).All(result)
}

//Find a Elements into the database
func (this *MgoConn) Find(collectionName string, search bson.M, result interface{}) error {

	return this.DB.C(collectionName).Find(search).All(result)

}

//Update a Element into the database
func (this *MgoConn) Update(collectionName string, search bson.M, change bson.M) error {
	return this.DB.C(collectionName).Update(search, change)

}

//Remove a element.
func (this *MgoConn) Delete(collectionName string, selector bson.M) error {
	return this.DB.C(collectionName).Remove(selector)
}
