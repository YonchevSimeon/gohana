package gohana

import (
	"crypto/tls"
	"database/sql"

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
}

func (*Instance) Disconnect() {
	Db.Close()
}
