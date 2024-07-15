package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.WithValue(context.Background(), "name", "Joe")

	log.Printf("name = %v", ctx.Value("name"))
	log.Printf("age = %v", ctx.Value("age"))
}
