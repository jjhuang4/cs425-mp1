package main

import (
	"cs425/mp1/logger"
	"cs425/mp1/query"
	"net"
	"net/rpc"
	"os"
)

// https://pkg.go.dev/net/rpc#ServeConn
// All exported methods must follow the format:
// func (t *T) MethodName(argType T1, replyType *T2) error
// (t *T) a method receiver indicating that the method is made to interact with the type T

// only the server knows the vm_id

func main() {
	// register type / method here
	queryObj := new(query.Query)
	rpc.Register(queryObj)
	curVM := os.Args[1]
	curIP := query.Vm_to_ip[curVM]
	listener, err := net.Listen("tcp", curIP)

	// vm_id, err := strconv.Atoi(curVM)
	// if err != nil {
	// 	panic(err) // do not surface logging before logger.Init() is called
	// }
	logger.Init()
	if err != nil {
		logger.Server.Error("Error occurred listening for connection:" + err.Error())
		return
	}

	logger.Server.Info("Listening on port: " + curIP)
	// fmt.Println("Listening on port", curIP)
	rpc.Accept(listener)
	defer listener.Close()
}
