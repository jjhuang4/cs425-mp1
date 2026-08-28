package main

import (
	"flag"
	"fmt"
	"net/rpc"
)

// https://pkg.go.dev/net/rpc#ServeConn
// All exported methods must follow the format:
// func (t *T) MethodName(argType T1, replyType *T2) error
// (t *T) a method receiver indicating that the method is made to interact with the type T

// type Dummy struct{}

// func (d *Dummy) Greet(arg *string, reply *string) error {
// 	if arg == nil {
// 		return errors.New("argument cannot be empty")
// 	}
// 	*reply = "Greetings"
// 	return nil
// }

func main() {
	// register type / method here
	var port int
	flag.IntVar(&port, "port", 8080, "Port to connect to server")
	flag.Parse()

	conn_path := fmt.Sprintf("localhost:%d", port)
	client, err := rpc.Dial("tcp", conn_path)
	if err != nil {
		fmt.Println("Error occurred connecting to server:", err)
		return
	}
	defer client.Close()

	var dummy_str = "fill string"
	var reply string

	err = client.Call("Dummy.Greet", dummy_str, &reply)
	if err != nil {
		fmt.Println("Error occurred calling server method with RPC:", err)
		return
	}
	fmt.Println("Reply from server:", reply)
}
