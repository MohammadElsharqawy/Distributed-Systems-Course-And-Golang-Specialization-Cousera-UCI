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

	log.Println(messageObj.User + " says " + messageObj.Message)
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
