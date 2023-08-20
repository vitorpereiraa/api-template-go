package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vitorpereira/api-template-go/dtos"
	"github.com/vitorpereira/api-template-go/services"
)

type todoListRouter struct {
	routes *gin.RouterGroup
	svc    *services.TodoListService
}

func InitializeTodoListRouter(svc *services.TodoListService, router *gin.RouterGroup) {
	routes := router.Group("/todoList")

	r := &todoListRouter{
		routes: routes,
		svc:    svc,
	}

	r.routes.GET("/", r.findTodoLists)
	r.routes.GET("/:id", r.findTodoListByID)
	r.routes.POST("/", r.createTodoList)
	r.routes.DELETE("/:id", r.deleteTodoListByID)
}

func (t *todoListRouter) findTodoLists(ctx *gin.Context) {
	todoLists := t.svc.FindTodoLists()
	ctx.IndentedJSON(http.StatusOK, todoLists)
}

func (t *todoListRouter) findTodoListByID(ctx *gin.Context) {
	id := ctx.Param("id")

	todoList, err := t.svc.FindTodoListByID(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo List was not found"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, todoList)
}

func (t *todoListRouter) createTodoList(ctx *gin.Context) {
	var todoListDto dtos.TodoListDTO

	if err := ctx.BindJSON(&todoListDto); err != nil {
		return
	}

	res, err := t.svc.CreateTodoList(todoListDto)

	if err != nil {
		return
	}

	ctx.IndentedJSON(http.StatusCreated, res)
}

func (t *todoListRouter) deleteTodoListByID(ctx *gin.Context) {
	id := ctx.Param("id")

	deletedTodoList, err := t.svc.DeleteTodoListByID(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo List was not found"})
		return
	}

	ctx.IndentedJSON(http.StatusOK, deletedTodoList)
}
