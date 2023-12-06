/*
TODO
there are two implementations
1) pooling
	- the client will dial the rpc of the coordinating server
	- the client will call the remote procedure on the server to send a message
	- the client can fetch all of the messages history from the server using remote procedure call


2) event-driven [BONUS]
	- a client starts by looking for a port to establish it's server on (like giving my phone number to my friends to call me)
	- a client can send a message through an infinite loop waiting for input text, this message will be broadcasted to other clients through an rpc call on the server
	- a client can also receive messages simultaneously using the GO keyword
	- so a client here is a server + a client at the same time
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"rpc_assign/commons"
)

func main() {
	rpc_client, err := rpc.Dial("tcp", commons.Get_server_address())

	if err != nil {
		log.Fatal(err)
	}
	in := bufio.NewReader(os.Stdin)

	var username string

	fmt.Println("Enter your nickname :)")
	usernameByte, _, _ := in.ReadLine()

	username = string(usernameByte)
	var reply string

	err = rpc_client.Call("Listener.Register", username, &reply)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Reply: %s", reply)

	for {

		line, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		message := commons.Message{
			User:    username,
			Message: string(line),
		}

		var reply bool
		err = rpc_client.Call("Listener.SendMessage", message, &reply)
		if err != nil {
			log.Fatal(err)
		}

		var reply2 []string
		err = rpc_client.Call("Listener.RefreshMyChat", username, &reply2)
		if err != nil {
			log.Fatal(err)
		}
		for i := range reply2 {
			log.Println(reply2[i])
		}
	}

}
