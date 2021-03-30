package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Printf("Starting. . .\n")

	go func() {
		handleRequests()
	}()

	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	for {
		sig := <-signalChannel
		switch sig {
		case os.Interrupt:
			os.Exit(1)
		case syscall.SIGTERM:
			os.Exit(2)
			return
		}
	}
}
