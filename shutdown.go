package shutdown

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func WaitForShutdown() bool {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	s, ok := <-c
	if !ok {
		return false
	}
	log.Printf("Got signal: %v", s)

	return true
}
