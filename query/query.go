// package query

// import (
// 	"fmt"
// 	"os"
// 	"regexp"
// 	"bufio"
// )

// // see: https://pkg.go.dev/regexp, https://pkg.go.dev/net/rpc for docs
// // use regexp functions Compile/MustCompile, Find/FindALL, MatchString

// func ReadSingle(pattern string, file_path string) ([]byte, success bool) {
// 	contents, err := os.ReadFile(file_path)
// 	if err != nil {
// 		return (nil, false)
// 	}
// }

// func Grep(pattern string, file string) ([]byte, success bool) {
// 	// implement
// }