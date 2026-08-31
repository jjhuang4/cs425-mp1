package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"net/rpc"
	"os/exec"
)

// https://pkg.go.dev/net/rpc#ServeConn
// All exported methods must follow the format:
// func (t *T) MethodName(argType T1, replyType *T2) error
// (t *T) a method receiver indicating that the method is made to interact with the type T

// for testing RPC, create dummy type
type Dummy struct{}

func (d *Dummy) Greet(arg *string, reply *string) error {
	if arg == nil {
		return errors.New("argument cannot be empty")
	}
	*reply = "Greetings"
	return nil
}

type GrepRequest struct {
	Args []string
}
type GrepResponse struct {
	Results string
}

// if broadcast as RPC, each server receives same request
func (g *GrepRequest) Grep(args *GrepRequest, reply *GrepResponse) error {
	logfile := "logs/machine.*.log"
	grepArgs := append(args.Args, logfile)
	fmt.Println("Grep args:", grepArgs)

	cmd := exec.Command("grep", grepArgs...) // exec Command func returns Cmd struct
	fmt.Println("Executing command:", cmd.String())

	output, err := cmd.CombinedOutput()
	reply.Results = string(output)
	return err
}

func main() {
	// register type / method here
	dummy := new(Dummy)
	rpc.Register(dummy)

	grep := new(GrepRequest)
	rpc.Register(grep)

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
