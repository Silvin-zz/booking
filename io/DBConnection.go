package io

import (
	"errors"
)

type DBConnection interface {
	GetTypeConnection() string
	Start()
}

func Conn(dbtype string, name string, host string, port int, username string, password string) (DBConnection, error) {

	switch dbtype {

	case "mongodb":
		//create a mongodb connection
		//return new(MgoConn), nil

		break
	}
	return nil, errors.New("Invalid database type ")
}
