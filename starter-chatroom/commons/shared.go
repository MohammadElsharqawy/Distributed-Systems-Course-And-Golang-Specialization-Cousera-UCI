package commons

/*
TODO
you may define any structs here to be used by the rpc
*/

type Message struct {
	User    string
	Message string
}

// hint: you will need to have the server address fixed between clients and the coordinating server
func Get_server_address() string {
	return "0.0.0.0:9989"
}
