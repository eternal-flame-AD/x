package debug

import (
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
)

// PrintStackOnSIGINT starts a new goroutine and prints stack trace on SIGINT
func PrintStackOnSIGINT() {
	c := make(chan os.Signal)
	go func() {
		<-c
		debug.PrintStack()
		os.Exit(2)
	}()
	signal.Notify(c, syscall.SIGINT)
}
