package main

import (
	"dbdesign"
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	var reply dbdesign.Row

	client, err := rpc.DialHTTP("tcp", "localhost:3000")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	// data := dbdesign.Row{
	// 	Age:       65,
	// 	FirstName: "Jay",
	// 	LastName:  "Leno",
	// 	Email:     "jay@leno.io",
	// }

	// err = client.Call("API.Add", data, &reply)
	// if err != nil {
	// 	log.Fatal("error occured during Adding a User: ", err)
	// }
	// fmt.Println("UserAdded: ", reply)

	// data := dbdesign.Row{
	// 	ID:        17,
	// 	Age:       65,
	// 	FirstName: "Conan",
	// 	LastName:  "Brian",
	// 	Email:     "Conan@Brian.io",
	// }

	// err = client.Call("API.Update", data, &reply)
	// if err != nil {
	// 	log.Fatal("error occured during Adding a User: ", err)
	// }
	// fmt.Println("UserUpdated: ", reply)

	// data := dbdesign.Row{
	// 	ID: 12,
	// }

	// err = client.Call("API.Delete", data, &reply)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("UserDeleted")

	data := dbdesign.Row{
		ID: 120,
	}

	err = client.Call("API.Find", data, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("UserAdded", reply)
}
