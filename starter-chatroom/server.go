/*

TODO
the server has either two implementations
1) pooling
	- every message sent to the server has to be stored in long list
	- a client may ask for this list or a slice of it to fetch the updates


2) event-driven [BONUS]
	- a server is more like a coordinator
	- the server waits for clients wanting to register themselves as listeners
	- a client sends a message by calling an rpc responsible for broadcasting, the client calls a function to loop on all registered clients and send his own message to each of them separately
	- the client on the other side has a server listening for messages being pushed
*/

package main

import (
	"log"
	"net"
	"net/rpc"
	"rpc_assign/commons"
)

type Listener int

var messageQueue map[string][]string

func (l *Listener) Register(username string, reply *string) error {
	*reply += "Welcome to our public chat, Write anything to broadcast!"

	messageQueue[username] = nil

	//notify all
	for userKey, _ := range messageQueue {
		messageQueue[userKey] = append(messageQueue[userKey], username+" connected to the chat.")
	}

	log.Printf("%s has joined the chat.\n", username)

	return nil
}

// RPC Send
func (l *Listener) SendMessage(messageObj commons.Message, reply *bool) error {
	for key, _ := range messageQueue {
		messageQueue[key] = append(messageQueue[key], messageObj.User+" says "+messageObj.Message)
	}
	*reply = true
	return nil
}

func (l *Listener) RefreshMyChat(username string, reply *[]string) error {
	*reply = messageQueue[username]
	messageQueue[username] = nil
	return nil
}

func main() {
	messageQueue = make(map[string][]string)

	addr, err := net.ResolveTCPAddr("tcp", commons.Get_server_address())
	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	listener := new(Listener)
	rpc.Register(listener)

	rpc.Accept(inbound)

}
