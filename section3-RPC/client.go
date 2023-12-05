package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"section3/packages/commons"
)

func check(rpc_client *rpc.Client, username string) {
	var reply []string
	err := rpc_client.Call("Listener.RefreshMyChat", username, &reply)
	if err != nil {
		log.Fatal(err)
	}
	for i := range reply {
		log.Println(reply[i])
	}

}

func send(rpc_client *rpc.Client, username string, line []byte) {

	message := commons.Message{
		User:    username,
		Message: string(line),
	}

	var reply bool
	err := rpc_client.Call("Listener.SendMessage", message, &reply)
	if err != nil {
		log.Fatal(err)
	}

}

func register(rpc_client *rpc.Client, usernameByte []byte) {
	var reply string
	err := rpc_client.Call("Listener.Register", string(usernameByte), &reply)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Reply: %s", reply)
}

func main() {
	rpc_client, err := rpc.Dial("tcp", "0.0.0.0:42580")

	if err != nil {
		log.Fatal(err)
	}
	in := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your nickname :)")
	usernameByte, _, _ := in.ReadLine()

	register(rpc_client, usernameByte)

	//fmt.Println("============Welcome=============")
	//fmt.Println("============Message RPC=============")

	for {

		//fmt.Println("Please, Write any message to be sent to the server")

		line, _, err := in.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		send(rpc_client, string(usernameByte), line)

		check(rpc_client, string(usernameByte))

	}

	/*
			fmt.Println("=====================window chat=================")

			for _, val := range messages {
				fmt.Println(val)
			}

			//////////////////////adding///////////////////////
			//
			//fmt.Println("Enter the first operand")
			//first_operand, _, err := in.ReadLine()
			//if err != nil {
			//	log.Fatal(err)
			//}
			//a, _ := strconv.Atoi(string(first_operand))
			//fmt.Println("Enter the second operand")
			//second_operand, _, err := in.ReadLine()
			//if err != nil {
			//	log.Fatal(err)
			//}
			//b, _ := strconv.Atoi(string(second_operand))
			//args := &commons.Args{a, b}
			//
			//var reply2 int
			//err = rpc_client.Call("Listener.Add", args, &reply2)
			//
			//fmt.Printf("Arith: %d+%d=%d\n", args.A, args.B, reply2)

			fmt.Println("================Thanks================")
		}
	*/
}
