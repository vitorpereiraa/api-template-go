package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitorpereira/api-template-go/services"
)

func Initialize() {
	r := gin.Default()

	r.GET("/todolist", getTodoLists)
	// r.POST("/todolist", createTodoLists)

	r.Run()
}

func getTodoLists(ctx *gin.Context) {
	todoLists := services.GetTodoLists()

	ctx.IndentedJSON(http.StatusOK, todoLists)
}

// func createTodoLists(ctx *gin.Context) {

// }
