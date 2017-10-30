package io

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

type MgoConn struct {
	DB   *mgo.Database
	Conn *mgo.Session
}

/**
* Create a MongoDB Connection
*
**/
func (this *MgoConn) New(url string, database string) (error, string) {

	//we can use a auth parameters in production
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{url},
		Database: beego.AppConfig.String(database),
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

/*
* Close a MongoDB Connection
 */
func (this *MgoConn) Close() {
	session.Close()
}
