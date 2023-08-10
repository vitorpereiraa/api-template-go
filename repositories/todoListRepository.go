package repositories

import (
	"github.com/vitorpereira/api-template-go/domain"
)

var todoListRepository []domain.TodoList

func init() {
	var items []domain.Item

	items = append(items, domain.NewItem("Do Homework", "Until Tomorrow"))

	var todoList domain.TodoList = domain.NewTodoList("School Todo List", items)

	todoListRepository = append(todoListRepository, todoList)
}

func GetTodoLists() []domain.TodoList {
	return todoListRepository
}
