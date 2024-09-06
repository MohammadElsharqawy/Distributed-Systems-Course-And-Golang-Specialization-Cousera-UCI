package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"rpc_assign/commons"
)

func register(rpc_client *rpc.Client, username string) {
	var reply string
	err := rpc_client.Call("Listener.Register", username, &reply)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Reply: %s", reply)

	return
}

func send(rpc_client *rpc.Client, message commons.Message) {
	var reply bool
	err := rpc_client.Call("Listener.SendMessage", message, &reply)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func refresh(rpc_client *rpc.Client, username string) {
	var reply2 []string
	err := rpc_client.Call("Listener.RefreshMyChat", username, &reply2)
	if err != nil {
		log.Fatal(err)
	}
	for i := range reply2 {
		log.Println(reply2[i])
	}
	return
}

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
	register(rpc_client, username)

	for {

		line, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		message := commons.Message{
			User:    username,
			Message: string(line),
		}

		send(rpc_client, message)

		refresh(rpc_client, username)
	}

}
