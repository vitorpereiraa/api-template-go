package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitorpereira/api-template-go/config"
	"github.com/vitorpereira/api-template-go/repositories"
	"github.com/vitorpereira/api-template-go/services"
)

//TODO: Add support for open api definition

func Initialize() {
	r := gin.Default()

	r.GET("/health", healthCheck)

	api := r.Group("/api/v1")

	repo := repositories.NewTodoListRepository()
	svc := services.NewTodoListService(&repo)

	InitializeTodoListRouter(svc, api)

	r.Run(config.Settings.SERVER_HOST + ":" + config.Settings.SERVER_PORT)
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
