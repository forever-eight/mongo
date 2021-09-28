package endpoint

import (
	"github.com/labstack/echo"

	"github.com/forever-eight/mongo.git/service"
)

type Endpoint struct {
	s *service.Service
}

func NewEndpoint(s *service.Service) *Endpoint {
	return &Endpoint{s: s}
}

func (e *Endpoint) InitRoutes() *echo.Echo {
	echoS := echo.New()
	// g := echoS.Group("/change")
	// g.GET("/", mainAdmin)
	echoS.POST("/change", e.change)
	echoS.POST("/delete", e.delete)
	echoS.POST("/add", e.add)
	echoS.GET("/find", e.find)
	return echoS
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
