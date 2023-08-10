package domain

import "github.com/google/uuid"

type TodoList struct {
	id          uuid.UUID
	description string
	items       []Item
}

func NewTodoList(descrition string, items []Item) TodoList {
	return TodoList{
		id:          uuid.New(),
		description: descrition,
		items:       items,
	}
}

func (t *TodoList) ID() string {
	return t.id.String()
}

func (t *TodoList) Description() string {
	return t.description
}

func (t *TodoList) SetDescription(description string) {
	t.description = description
}

func (t *TodoList) Items() []Item {
	return t.items
}

func (t *TodoList) SetItems(items []Item) {
	t.items = items
}
