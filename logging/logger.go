package logger

import (
	"os"
	"fmt"
)

func LogFile(contents string, vm int) (success bool) {
	file, res := os.Create(fmt.Sprintf("machine.%d.log", vm))
	if res != nil {
		return false
	}
	file.WriteString(contents)
	defer file.Close()
	return true
}