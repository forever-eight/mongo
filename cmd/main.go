package main

import (
	"context"
	"fmt"
	"log"

	"github.com/forever-eight/mongo.git/ds"
	"github.com/forever-eight/mongo.git/repository"
)

const (
	dbUri = "mongodb://0.0.0.0:27017"
)

func main() {
	ctx := context.Background()

	// Инициализация репозитория
	rep, err := repository.New(ctx, dbUri)
	if err != nil {
		log.Fatal("Can't init repository:", err)
	}
	project := ds.Projects{
		Title: "МАИ",
		Channels: ds.Channels{
			Vk: []ds.VkConfig{ds.VkConfig{Token: "12345"}, ds.VkConfig{Token: "6555"}},
		},
	}
	err = rep.AddProject(ctx, &project)
	if err != nil {
		log.Println("Add error", err)
	}

	res, err := rep.FindProjectByID(ctx, "МТУСИ")
	if err != nil {
		log.Println("Find error", err)
	}

	for n := 0; n < len(res); n++ {
		fmt.Println(res[n])
	}
}
