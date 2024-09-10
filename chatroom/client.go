

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
