package main

import (
	"cs425/mp1/query"
	"errors"
	"flag"
	"fmt"
	"net"
	"net/rpc"
)

// https://pkg.go.dev/net/rpc#ServeConn
// All exported methods must follow the format:
// func (t *T) MethodName(argType T1, replyType *T2) error
// (t *T) a method receiver indicating that the method is made to interact with the type T

type Dummy struct{}

func (d *Dummy) Greet(arg *string, reply *string) error {
	if arg == nil {
		return errors.New("argument cannot be empty")
	}
	*reply = "Greetings"
	return nil
}

func main() {
	// register type / method here
	// dummy := new(Dummy)
	// rpc.Register(dummy)

	// var port int
	// flag.IntVar(&port, "port", 8080, "Port to listen on")
	// flag.Parse()

	// conn_path := fmt.Sprintf(":%d", port)
	// listener, err := net.Listen("tcp", conn_path)
	// if err != nil {
	// 	fmt.Println("Error occurred listening for connection:", err)
	// 	return
	// }
	// defer listener.Close()

	// fmt.Println("Listening on port", port)
	// rpc.Accept(listener)

	query := new(query.Query)
	rpc.Register(query)
	var port int
	flag.IntVar(&port, "port", 8080, "Port to listen on")
	flag.Parse()
	conn_path := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", conn_path)
	if err != nil {
		fmt.Println("Error occurred listening for connection:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Listening on port", port)
	rpc.Accept(listener)

}
