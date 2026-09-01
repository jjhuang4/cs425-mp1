package main

import (
	"cs425/mp1/logger"
	"cs425/mp1/query"
	"net/rpc"
	"os"
	"sync"
)

// https://pkg.go.dev/net/rpc#ServeConn
// All exported methods must follow the format:
// func (t *T) MethodName(argType T1, replyType *T2) error
// (t *T) a method receiver indicating that the method is made to interact with the type T

// var vm_to_ip = map[string]string{"vm1": "127.0.0.1", "vm2": "127.0.0.1"}

type GrepArgs struct {
	Flags   string
	Pattern string
	File    *string
}

func grepCall(vm, ip string) error {

	// fmt.Printf("Processing VM at: %s", ip)
	logger.Client.Info("Processing VM at: " + ip)
	client, err := rpc.Dial("tcp", string(ip))
	if err != nil {
		// fmt.Println("Error occurred connecting to server:", err)
		logger.Client.Error("Error occurred connecting to server:" + err.Error())
		return err
	}
	defer client.Close()
	var reply query.Reply

	grepArgs := &query.GrepArgs{
		Flags:   os.Args[2],
		Pattern: os.Args[3],
		File:    &os.Args[4],
	}
	// fmt.Printf("\nArgs provided to Grep command \nFlags: %s\nPattern: %s\nFilepath: %s",
	// 	grepArgs.Flags, grepArgs.Pattern, *grepArgs.File)
	logger.Client.Info("Args provided to Grep command \nFlags: " + grepArgs.Flags + "\nPattern: " + grepArgs.Pattern + "\nFilepath: " + *grepArgs.File)
	err = client.Call("Query.Grep", grepArgs, &reply)
	// err = client.Call("Query.Grep", os.Args[3:], &reply)

	if err != nil {
		// fmt.Println("Error occurred calling server method with RPC:", err)
		logger.Client.Error("Error occurred calling server method with RPC:" + err.Error())
		return err
	}
	// fmt.Print("\nReply from server: \n", string(reply.Reply))
	logger.Client.Info("Reply from server: " + string(reply.Reply))
	return nil
}

func main() {
	logger.Init()

	if len(os.Args) == 1 {
		// fmt.Println("No args")
		logger.Client.Warn("No args provided")
		return
	}
	current_vm := os.Args[2]
	// vm_name := os.Args[1]
	// vm_to_ip[vm_name]

	// https://gobyexample.com/waitgroups
	var wg sync.WaitGroup

	for vm, ip := range query.Vm_to_ip {

		if vm == current_vm {
			// fmt.Println("Skipping current VM:", vm)
			logger.Client.Info("Skipping current VM:" + vm)
			continue
		}

		wg.Go(func() {
			grepCall(vm, ip)
		})
	}
	wg.Wait()
}
