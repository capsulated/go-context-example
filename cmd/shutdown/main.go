package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	doWork(ctx)
}

func doWork(ctx context.Context) {
	newCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	log.Println("starting working...")

	for {
		select {
		case <-newCtx.Done():
			log.Printf("ctx done: %v", ctx.Err())
			return
		default:
			log.Println("working...")
			time.Sleep(1 * time.Second)
		}
	}
}
