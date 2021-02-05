package main

import (
	"fmt"
)


func main()  {
	fmt.Println("Entrer votre Nom: ") 
    var lastName string 
	fmt.Scanln(&lastName)
	
    fmt.Println("Entrer votre prenom: ") 
    var firstName string 
    fmt.Scanln(&firstName) 
  
    fmt.Println("Entrer votre date de naissance: ") 
    var birthdayDate string 
	fmt.Scanln(&birthdayDate) 


	var myClient *client
	myClient = newClient(lastName, firstName, birthdayDate)

	fmt.Println("Voici votre age: ", myClient.age)
}