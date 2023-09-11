package services

import "github.com/vitorpereira/api-template-go/dtos"

type ITodoListService interface {
	FindTodoLists() []dtos.TodoListDTO
	FindTodoListByName(name string) (dtos.TodoListDTO, error)
	CreateTodoList(todoListDTO dtos.TodoListDTO) (dtos.TodoListDTO, error)
	DeleteTodoListByName(name string) (dtos.TodoListDTO, error)
}