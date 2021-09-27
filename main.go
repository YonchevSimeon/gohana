package main

import "github.com/YonchevSimeon/gohana/gohana"

func main() {

	gohana := &gohana.Instance{}

	gohana.Connect("5aa93e26-6c8c-4f17-8145-82035219d7bb.hana.trial-us10.hanacloud.ondemand.com", "443", "DBADMIN", "Qwerty753951")
	defer gohana.Disconnect()
}