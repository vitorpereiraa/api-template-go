package dtos

type TodoListDTO struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Items       []ItemDTO `json:"items"`
}
