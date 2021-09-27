package gohana

import (
	"crypto/tls"
	"database/sql"
	"fmt"

	"github.com/SAP/go-hdb/driver"
)

var Db *sql.DB

type Instance struct {
}

func (*Instance) Connect(host, port, username, password string) {
	c := driver.NewBasicAuthConnector(host+":"+port, username, password)

	tlsConfig := tls.Config{
		InsecureSkipVerify: false,
		ServerName:         host,
	}

	c.SetTLSConfig(&tlsConfig)
	Db = sql.OpenDB(c)
	if err := Db.Ping(); err != nil {
		fmt.Println("Connection interrupted! Please check your input parameters.")	
	}else{
		fmt.Println("Successfully connected!")
	}
}

func (*Instance) Disconnect() {
	Db.Close()
	fmt.Println("Successfully disconnected!")
}
