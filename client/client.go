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

	addUser := dbdesign.Row{
		Age:       65,
		FirstName: "Jay",
		LastName:  "Leno",
		Email:     "jay@leno.io",
	}

	client.Call("API.Add", addUser, &reply)

	fmt.Println("UserAdded: ", reply)
}
