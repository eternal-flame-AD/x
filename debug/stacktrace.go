package debug

import (
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
)

func PrintStackOnSIGINT() {
	c := make(chan os.Signal)
	go func() {
		<-c
		debug.PrintStack()
		os.Exit(2)
	}()
	signal.Notify(c, syscall.SIGINT)
}
