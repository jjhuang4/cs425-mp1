package client

import (
	"cs425/mp1/query"
	"fmt"
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

// Helper function that accepts vm and its ip address as strings. Utilizes GO net/RPC
// to initiate client-server interaction between the current and target vm. Returns the
// output of the user-provided grep command from the server.
func GrepCall(vm, flags, pattern, file string) (string, error) {

	ip := query.VM_to_IP[vm]
	fmt.Printf("Processing VM at: %s", ip)
	client, err := rpc.Dial("tcp", string(ip))
	if err != nil {
		fmt.Println("Error occurred connecting to server:", err)
		return "", err
	}
	defer client.Close()
	var reply query.Reply

	grepArgs := &query.GrepArgs{
		Flags:   flags,
		Pattern: pattern,
		File:    &file,
	}

	fmt.Printf("\nArgs provided to Grep command \nFlags: %s\nPattern: %s\nFilepath: %s",
		grepArgs.Flags, grepArgs.Pattern, *grepArgs.File)

	err = client.Call("Query.Grep", grepArgs, &reply)

	if err != nil {
		fmt.Println("Error occurred calling server method with RPC:", err)
		return "", err
	}
	fmt.Print("\nReply from server: \n", string(reply.Reply))
	return string(reply.Reply), nil
}

func FileWriteCall(testLog, vm, file string) (string, error) {
	ip := query.VM_to_IP[vm]
	client, err := rpc.Dial("tcp", string(ip))
	if err != nil {
		fmt.Println("Error occurred connecting to server:", err)
		return "", err
	}

	defer client.Close()

	//Unused
	var reply query.Reply

	fileWriteArgs := query.FileWriteArgs{
		TestLog: testLog,
		File:    &file,
		VM:      vm,
	}

	fmt.Printf("\nArgs provided to fileWriteCall \nFlags: %s\nPattern: %s\nFilepath: %s",
		fileWriteArgs.TestLog, *fileWriteArgs.File, fileWriteArgs.VM)

	err = client.Call("Query.FileWrite", fileWriteArgs, &reply)
	if err != nil {
		fmt.Println("Error occurred calling FileWrite with RPC:", err)
		return "", err
	}

	return string(reply.Reply), nil

}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("No args")
		return
	}
	currentVM := os.Args[1]
	flags := os.Args[2]
	pattern := os.Args[3]
	file := os.Args[4]

	// https://gobyexample.com/waitgroups
	fmt.Println(currentVM)

	var wg sync.WaitGroup
	for vm := range query.VM_to_IP {

		// if vm == currentVM {
		// 	fmt.Println("Skipping current VM:", vm)
		// 	continue
		// }

		wg.Go(func() {
			_, err := GrepCall(vm, flags, pattern, file)
			if err != nil {
				fmt.Printf("Error when calling %s: %s", vm, err)
			}
			// fmt.print(reply)
		})
	}
	wg.Wait()
}
