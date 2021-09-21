package main

import (
	"context"
	"log"

	"github.com/forever-eight/mongo.git"
	"github.com/forever-eight/mongo.git/endpoint"
	"github.com/forever-eight/mongo.git/repository"
	"github.com/forever-eight/mongo.git/service"
)

const (
	dbUri = "mongodb://0.0.0.0:27017"
)

func main() {
	ctx := context.Background()

	// Инициализация репозитория, сервиса и эндпойнта
	rep, err := repository.NewRepos(ctx, dbUri)
	if err != nil {
		log.Fatal("Can't init repository:", err)
	}
	services := service.NewService(rep)
	endpoints := endpoint.NewEndpoint(services)
	// todo echo подключить и в эндпойнт
	srv := new(mongo.Server)
	err = srv.Run(endpoints.InitRoutes())
	if err != nil {
		log.Println("error occurred while running http server ", err)
	}

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

	/*project := ds.Project{
		Title: "Машенька",
		Channels: ds.Channels{
			Vk: []ds.VkConfig{ds.VkConfig{Token: "213"}, ds.VkConfig{Token: "55555"}},
			Tg: []ds.TgConfig{ds.TgConfig{Token: "1234"}},
		},
	}

	err = rep.ChangeProject(ctx, "613271e6aaadd5d13e525940", &project)

	if err != nil {
		log.Println("change error", err)
	}*/

}
