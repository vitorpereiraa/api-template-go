package dtos

import "github.com/vitorpereira/api-template-go/domain"

func TodoListToDto(todoList domain.TodoList) TodoListDTO {
	return TodoListDTO{
		Description: todoList.Description(),
		Items:       ItemListToDtoList(todoList.Items()),
	}
}

func TodoListListToDtoList(todoListList []domain.TodoList) (todoListDtos []TodoListDTO) {
	for _, todoList := range todoListList {
		todoListDtos = append(todoListDtos, TodoListToDto(todoList))
	}
	return
}
