package main

import {
	fmt
	client
}

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
	
	layout := "14/02/2010"
	t, err := time.Parse(layout, birthdayDate)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)


	var client myClient*
	//myClient = client.newClient(,,,,)
}