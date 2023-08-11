package dtos

import "github.com/vitorpereira/api-template-go/domain"

func TodoListToDto(todoList domain.TodoList) TodoListDTO {
	return TodoListDTO{
		ID:          todoList.ID(),
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

func TodoListDtoToTodoList(todoListDTO TodoListDTO) domain.TodoList {
	return domain.NewTodoList(todoListDTO.Description, ItemDtoListToItemList(todoListDTO.Items))
}
