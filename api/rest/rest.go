package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vitorpereira/api-template-go/config"
	"github.com/vitorpereira/api-template-go/repositories"
	"github.com/vitorpereira/api-template-go/services"
)

func Initialize() {
	r := gin.Default()

	api := r.Group("/api/v1")

	repo := repositories.NewTodoListRepository()
	svc := services.NewTodoListService(&repo)

	InitializeTodoListRouter(svc, api)

	r.Run(config.Settings.SERVER_HOST + ":" + config.Settings.SERVER_PORT)
}
