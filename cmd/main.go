package main

import (
	"context"
	"encoding/json"
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

	// Создать проект
	/*project := ds.Project{
		ID:    primitive.NewObjectID(),
		Title: "Мария",
		Channels: ds.Channels{
			Vk:   []ds.VkConfig{ds.VkConfig{Token: "213"}, ds.VkConfig{Token: "ьувтвтв111"}},
			Tg:   []ds.TgConfig{ds.TgConfig{Token: "1234"}},
		},
	}
	err = rep.AddProject(ctx, &project)
	if err != nil {
		log.Println("Add error", err)
	}*/

	// Найти проект

	/*found, err := rep.FindProjectByID(ctx, "612e7947290806db4c6490f6")
	if err != nil {
		log.Println("find error", err)
	}

	foundJSON, err := json.MarshalIndent(found, "", "\t")
	if err != nil {
		log.Println("json error", err)
	}

	fmt.Print(string(foundJSON))*/

	// Удаление проекта

	/*	err = rep.DeleteProject(ctx, "612e7947290806db4c6490f6")
		if err != nil {
			log.Println("delete error", err)
		}*/

	// Изменение проекта
	project := ds.Project{
		Title: "Мария",
		Channels: ds.Channels{
			Vk: []ds.VkConfig{ds.VkConfig{Token: "213"}, ds.VkConfig{Token: "ьувтвтв111"}},
			Tg: []ds.TgConfig{ds.TgConfig{Token: "1234"}},
		},
	}

	changed, err := rep.RemakeProject(ctx, "613271e6aaadd5d13e525940", &project)

	if err != nil {
		log.Println("change error", err)
	}

	changedJSON, err := json.MarshalIndent(changed, "", "\t")
	if err != nil {
		log.Println("json error", err)
	}
	fmt.Print(string(changedJSON))
}
