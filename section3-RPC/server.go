package main

import (
	"log"
	"net"
	"net/rpc"
	"section3/packages/commons"
)

type Listener int

var messageQueue map[string][]string
var users []string

func (l *Listener) Register(username string, reply *string) error {
	*reply += "Welcome to our public chat, Write anything to broadcast!"

	users = append(users, username)
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

	log.Printf(messageObj.User + " says " + messageObj.Message)
	*reply = true
	return nil
}

func (l *Listener) RefreshMyChat(username string, reply *[]string) error {
	*reply = messageQueue[username]
	messageQueue[username] = nil
	return nil
}

//func (l *Listener) Add(args *commons.Args, reply *int) error {
//	*reply = (*args).A + (*args).B
//	fmt.Printf("Arith: %d+%d done on the server\n", args.A, args.B)
//	return nil
//}

func main() {

	//listener := new(Listener)
	//var reply bool
	//listener.GetLine([]byte("hi"), &reply)

	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:42580")
	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	messageQueue = make(map[string][]string)
	listener := new(Listener)
	rpc.Register(listener)

	rpc.Accept(inbound)

}
