package main

import (
	"context"
	"log"

	"github.com/forever-eight/mongo.git/repository"
)

const (
	dbUri = "mongodb://0.0.0.0:27017"
)

func main() {
	ctx := context.Background()

	// Инициализация репозитория
	_, err := repository.New(ctx, dbUri)
	if err != nil {
		log.Fatal("Can't init repository:", err)
	}
	//ctx = repository.CreateContext(ctx, repos)
}
