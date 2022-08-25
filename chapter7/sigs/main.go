package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(
		signals,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	go func() {
		switch <-signals {
		case syscall.SIGINT, syscall.SIGTERM:
			cleanup()
			os.Exit(1)
		case syscall.SIGQUIT:
			cleanup()
			panic("SIGQUIT called")
		}
	}()
}

func cleanup() {
	fmt.Println("hogeghgoe")
}
