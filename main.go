package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"p2p-network-simulator/http"
)

func main() {
	ctx := context.Background()

	httpServer := http.NewHTTPServer()
	httpServer.Start()

	channel := make(chan os.Signal, 1)

	signal.Notify(channel, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	<-channel

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	httpServer.Shutdown(ctx)

	os.Exit(0)
}
