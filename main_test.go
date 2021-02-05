package main

import (
	"testing"
)

func TestNewClient(t *testing.T) {
    var lastName string 
    var firstName string 
    var birthdayDate string 

	lastName = "Dufresne"
	firstName = "Enzo"
	birthdayDate = "01/01/97"

	var myClient *client
	myClient = newClient(lastName, firstName, birthdayDate)
	if myClient == nil {
		t.Error("Le client n'a pas été initialiser")
	}
	if myClient.lastName != "Dufresne" {
		t.Error("LastName false")
	}
}