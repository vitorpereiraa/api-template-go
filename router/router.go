package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitorpereira/api-template-go/dtos"
	"github.com/vitorpereira/api-template-go/services"
)

func Initialize() {
	r := gin.Default()

	r.GET("/todolist", findTodoLists)
	r.GET("/todolist/:id", findTodoListByID)
	r.POST("/todolist", createTodoList)
	r.DELETE("/todolist/:id", deleteTodoListByID)

	r.Run()
}

func findTodoLists(ctx *gin.Context) {
	todoLists := services.FindTodoLists()
	ctx.IndentedJSON(http.StatusOK, todoLists)
}

func findTodoListByID(ctx *gin.Context) {
	id := ctx.Param("id")

	todoList, err := services.FindTodoListByID(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo List was not found"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, todoList)
}

func createTodoList(ctx *gin.Context) {
	var todoListDto dtos.TodoListDTO

	if err := ctx.BindJSON(&todoListDto); err != nil {
		return
	}

	res := services.CreateTodoList(todoListDto)

	ctx.IndentedJSON(http.StatusCreated, res)
}

func deleteTodoListByID(ctx *gin.Context) {
	id := ctx.Param("id")

	deletedTodoList, err := services.DeleteTodoListByID(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo List was not found"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, deletedTodoList)
}
