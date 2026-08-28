package logger

import (
	"fmt"
	"os"
)

func LogFile(contents string, vm int) (success bool) {
	file, err := os.Create(fmt.Sprintf("machine.%d.log", vm))
	if err != nil {
		return false
	}
	file.WriteString(contents)
	defer file.Close()
	return true
}
