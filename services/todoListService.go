package services

import (
	"github.com/vitorpereira/api-template-go/dtos"
	"github.com/vitorpereira/api-template-go/repositories"
)

func GetTodoLists() []dtos.TodoListDTO {

	todoLists := repositories.GetTodoLists()

	return dtos.TodoListListToDtoList(todoLists)
}
