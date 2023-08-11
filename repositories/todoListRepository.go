package repositories

import (
	"errors"

	"github.com/vitorpereira/api-template-go/domain"
)

var todoListRepository []domain.TodoList

func init() {
	var items []domain.Item

	items = append(items, domain.NewItem("Do Homework", "Until Tomorrow"))

	var todoList domain.TodoList = domain.NewTodoList("School Todo List", items)

	todoListRepository = append(todoListRepository, todoList)
}

func FindTodoLists() []domain.TodoList {
	return todoListRepository
}

func FindTodoListByID(id string) (domain.TodoList, error) {
	for i := 0; i < len(todoListRepository); i++ {
		if todoListRepository[i].ID() == id {
			return todoListRepository[i], nil
		}
	}
	return domain.TodoList{}, errors.New("Todo List Not Found")
}

func CreateTodoList(todoList domain.TodoList) domain.TodoList {
	todoListRepository = append(todoListRepository, todoList)
	return todoList
}

func DeleteTodoListByID(id string) (domain.TodoList, error) {
	var todoList domain.TodoList

	for i := 0; i < len(todoListRepository); i++ {
		if todoListRepository[i].ID() == id {
			todoListRepository, todoList = remove(todoListRepository, i)
			return todoList, nil
		}
	}

	return todoList, errors.New("Todo List Not Found")
}

func remove(s []domain.TodoList, i int) ([]domain.TodoList, domain.TodoList) {
	s[i] = s[len(s)-1]
	return s[:len(s)-1], s[i]
}
