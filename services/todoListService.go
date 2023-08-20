package services

import (
	"github.com/vitorpereira/api-template-go/dtos"
	"github.com/vitorpereira/api-template-go/repositories"
)

type TodoListService struct {
	repo repositories.ITodoListRepository
}

func NewTodoListService(todoListRepo repositories.ITodoListRepository) *TodoListService {
	return &TodoListService{
		repo: todoListRepo,
	}
}

func (s TodoListService) FindTodoLists() []dtos.TodoListDTO {
	todoLists := s.repo.FindTodoLists()
	return dtos.TodoListListToDtoList(todoLists)
}

func (s TodoListService) FindTodoListByID(id string) (dtos.TodoListDTO, error) {
	todoList, err := s.repo.FindTodoListByID(id)
	return dtos.TodoListToDto(todoList), err
}

func (s TodoListService) CreateTodoList(todoListDTO dtos.TodoListDTO) (dtos.TodoListDTO, error) {
	todoList := dtos.TodoListDtoToTodoList(todoListDTO)
	createdTodoList, err := s.repo.CreateTodoList(todoList)
	return dtos.TodoListToDto(createdTodoList), err
}

func (s TodoListService) DeleteTodoListByID(id string) (dtos.TodoListDTO, error) {
	todoList, err := s.repo.DeleteTodoListByID(id)
	return dtos.TodoListToDto(todoList), err
}
