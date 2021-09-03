package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

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
	/*	project := ds.Projects{
			Title: "МАИ",
			Channels: ds.Channels{
				Vk: []ds.VkConfig{ds.VkConfig{Token: "12345"}, ds.VkConfig{Token: "6555"}},
			},
		}
		err = rep.AddProject(ctx, &project)
		if err != nil {
			log.Println("Add error", err)
		}*/

	found, err := rep.FindProjectByID(ctx, "612e7947290806db4c6490f6")
	if err != nil {
		log.Println("find error", err)
	}

	foundJSON, err := json.MarshalIndent(found, "", "\t")
	if err != nil {
		log.Println("json error", err)
	}

	fmt.Print(string(foundJSON))

}
