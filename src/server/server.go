package main

import (
	"cs425/mp1/query"
	"fmt"
	"net"
	"net/rpc"
	"os"
)

// https://pkg.go.dev/net/rpc#ServeConn
// All exported methods must follow the format:
// func (t *T) MethodName(argType T1, replyType *T2) error
// (t *T) a method receiver indicating that the method is made to interact with the type T

func main() {
	// register type / method here

	queryObj := new(query.Query)
	rpc.Register(queryObj)
	curVM := os.Args[1]
	curIP := query.Vm_to_ip[curVM]
	listener, err := net.Listen("tcp", curIP)
	if err != nil {
		fmt.Println("Error occurred listening for connection:", err)
		return
	}

	rpc.Accept(listener)
	fmt.Println("Listening on port", curIP)
	defer listener.Close()
}
