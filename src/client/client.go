package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/rpc"
	"os"
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

// Purpose of this client is to act as the primary interface for greping servers
// Server is persistent, client stays alive for duration of distributed grep call
// Client sends request to all known active servers through RPC, collects reponses from each server

type GrepRequest struct {
	Args []string
}
type GrepResponse struct {
	Results string
}

func main() {
	// from client, access all known servers through servers.config file, broadcast with RPC
	file, err := os.Open("servers.txt") // list known server ports
	if err != nil {
		fmt.Println("Error occurred opening servers.txt file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		server := scanner.Text()

		fmt.Println(i, server)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading servers.txt:", err)
	}

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
		fmt.Println("Error occurred calling server method with dummy RPC:", err)
		return
	}
	fmt.Println("Reply from server:", reply)

	var grep_args = []string{"-i"}
	var grep_reply GrepResponse
	err = client.Call("GrepRequest.Grep", &GrepRequest{Args: grep_args}, &grep_reply)
	if err != nil {
		fmt.Println("Error occurred calling server method with grep RPC:", err)
		return
	}

	fmt.Println("Reply from server:", grep_reply.Results)
}
