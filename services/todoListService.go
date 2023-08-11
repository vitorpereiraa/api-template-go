package services

import (
	"github.com/vitorpereira/api-template-go/dtos"
	"github.com/vitorpereira/api-template-go/repositories"
)

func FindTodoLists() []dtos.TodoListDTO {
	todoLists := repositories.FindTodoLists()
	return dtos.TodoListListToDtoList(todoLists)
}

func FindTodoListByID(id string) (dtos.TodoListDTO, error) {
	todoList, err := repositories.FindTodoListByID(id)
	return dtos.TodoListToDto(todoList), err
}

func CreateTodoList(todoListDTO dtos.TodoListDTO) dtos.TodoListDTO {
	todoList := dtos.TodoListDtoToTodoList(todoListDTO)
	createdTodoList := repositories.CreateTodoList(todoList)
	return dtos.TodoListToDto(createdTodoList)
}

func DeleteTodoListByID(id string) (dtos.TodoListDTO, error) {
	todoList, err := repositories.DeleteTodoListByID(id)
	return dtos.TodoListToDto(todoList), err
}
