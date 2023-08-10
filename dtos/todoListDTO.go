package dtos

type TodoListDTO struct {
	Description string    `json:"description"`
	Items       []ItemDTO `json:"items"`
}
