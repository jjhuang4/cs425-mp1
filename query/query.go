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
type Reply struct {
	Reply []byte
}

func (query *Query) Grep(args []string, reply *Reply) error {

	fmt.Println(args)
	cmd := exec.Command("grep", args...) //variadic unpack
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("error when grep: %v", err)
		return err
	}
	reply.Reply = out
	return nil
}
