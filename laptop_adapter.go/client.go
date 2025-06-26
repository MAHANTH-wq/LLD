package main

import "fmt"

type client struct {
}

func (c *client) insertLightningChargeIntoComputer(com computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.insertIntoLightningPort()
}
