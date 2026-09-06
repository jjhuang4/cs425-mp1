package query

// see: https://pkg.go.dev/regexp, https://pkg.go.dev/net/rpc for docs
// use regexp functions Compile/MustCompile, Find/FindALL, MatchString

//	func ReadSingle(pattern string, file_path string) ([]byte, success bool) {
//		contents, err := os.ReadFile(file_path)
//		if err != nil {
//			return (nil, false)
//		}
//	}
import (
	"fmt"
	"os"
	"os/exec"
)

type Query struct{}

//	type Args struct {
//		Flags   string
//		Pattern string
//		File    *string
//	}

type GrepArgs struct {
	Flags   string
	Pattern string
	File    *string
}

type FileWriteArgs struct {
	TestLog string
	File    *string
	VM      string
}

type Reply struct {
	Reply []byte
}

var VM_to_IP = map[string]string{"vm1": "127.0.0.1:8080", "vm2": "127.0.0.1:8001", "vm3": "127.0.0.1:4001"}

// func (query *Query) Grep(args []string, reply *Reply) error {
func (query *Query) Grep(args *GrepArgs, reply *Reply) error {
	fmt.Println(args)
	// cmd := exec.Command("grep", args...)
	cmd := exec.Command("grep", args.Flags, args.Pattern)

	if args.File != nil {
		cmd.Args = append(cmd.Args, *args.File)
	}
	out, err := cmd.Output()

	if err != nil {
		fmt.Printf("error when grep: %v", err)
		return err
	}
	reply.Reply = out
	return nil
}

// func LogFile(contents, vm string) (success error) {
// 	file, err := os.Create(fmt.Sprintf("machine.%s.log", vm))
// 	if err != nil {
// 		return err
// 	}
// 	file.WriteString(contents)
// 	defer file.Close()
// 	return nil
// }

// Function that accepts file to write to, string to write to, and vm
func (query *Query) FileWrite(args *FileWriteArgs, reply *Reply) error {

	file, err := os.Create(fmt.Sprintf(*args.File, args.VM))
	if err != nil {
		return err
	}
	file.WriteString(args.TestLog)
	defer file.Close()
	return nil
}

// refactor grepCall to accept os.Args parameters
