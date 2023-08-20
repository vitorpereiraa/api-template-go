package repositories

import "github.com/vitorpereira/api-template-go/domain"

type ITodoListRepository interface {
	FindTodoLists() []domain.TodoList
	FindTodoListByID(id string) (domain.TodoList, error)
	CreateTodoList(todoList domain.TodoList) (domain.TodoList, error)
	DeleteTodoListByID(id string) (domain.TodoList, error)
}
