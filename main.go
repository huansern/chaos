package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	urls := os.Args[1:]
	if len(urls) < 1 {
		// print help
		return
	}

	if err := verifyURL(urls); err != nil {
		log.Fatal(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	chaos(ctx, urls)
	stop()
}
