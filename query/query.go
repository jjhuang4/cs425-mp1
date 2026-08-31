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
type Reply struct {
	Reply []byte
}

var Vm_to_ip = map[string]string{"vm1": "127.0.0.1:8080", "vm2": "127.0.0.1:8080"}

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
