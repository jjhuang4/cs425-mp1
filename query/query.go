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
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
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

// var Vm_to_ip = map[string]string{"vm1": "127.0.0.1:8080", "vm2": "127.0.0.1:8080"}

// NOTE: direct stdout here, instead of passing through logging file
func getAvailableServers(file_path string) (map[string]string, error) {
	file, err := os.OpenFile(file_path, os.O_RDONLY, 0)
	if err != nil {
		fmt.Printf("Error occurred opening file listing available servers: %s, error: %s", file_path, err.Error())
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	vm_to_ip_mapping := make(map[string]string) // create str-to-str vm/ip mapping
	for scanner.Scan() {                        // Go's for loop acts as a while loop while scanner.Scan() returns true
		var line string = scanner.Text()
		line_parts := strings.Split(line, ",")
		if len(line_parts) != 2 {
			return nil, fmt.Errorf("Error occured parsing line in servers file: %d comma separated parts found, expected 2", len(line_parts))
		}
		vm_to_ip_mapping[line_parts[0]] = line_parts[1]
	}
	if err := scanner.Err(); err != nil { // if scanner.Err() produces error, return nothing and throw error
		return nil, fmt.Errorf("Error occurred scanning file: %s", err.Error())
	}
	return vm_to_ip_mapping, nil
}

var Vm_to_ip, _ = getAvailableServers("servers.txt")

// func (query *Query) Grep(args []string, reply *Reply) error {
func (query *Query) Grep(args *GrepArgs, reply *Reply) error {
	// cmd := exec.Command("grep", args...)
	cmd := exec.Command("grep", args.Flags, args.Pattern)

	if args.File != nil {
		cmd.Args = append(cmd.Args, *args.File)
	}
	out, err := cmd.Output()

	if err != nil {
		return err
	}
	reply.Reply = out
	return nil
}
