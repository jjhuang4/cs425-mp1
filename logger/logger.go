package logger

import (
	"fmt"
	"log/slog"
	"os"
)

// declare supported logger groups -- only these will show up in the logs
// logging should occur ONLY in these layers
var Server *slog.Logger
var Client *slog.Logger
var RPC *slog.Logger

// solely for testing, not meant to be called in prod or interact with other module code
func LogFile(contents string, vm int) (success bool) {
	file, err := os.Create(fmt.Sprintf("machine.%d.log", vm))
	if err != nil {
		return false
	}
	file.WriteString(contents)
	defer file.Close()
	return true
}

func Init() error {
	// init vm_id sometime after
	var vm_id int = 1

	// open log file in write mode, appending without overwriting existing logs, since this will be used cross-module
	// also creates file if not exist yet
	file, err := os.OpenFile(fmt.Sprintf("logs/machine.%d.log", vm_id),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	handler := slog.NewTextHandler(file, opts)
	root_log := slog.New(handler)
	slog.SetDefault(root_log)

	Server = slog.Default().WithGroup("server").With("component", "SERVER")
	Client = slog.Default().WithGroup("client").With("component", "CLIENT")
	RPC = slog.Default().WithGroup("rpc").With("component", "RPC")

	return nil
}
