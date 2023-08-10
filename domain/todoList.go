package domain

type TodoList struct {
	description string
	items       []Item
}

func NewTodoList(descrition string, items []Item) TodoList {
	return TodoList{
		description: descrition,
		items:       items,
	}
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
