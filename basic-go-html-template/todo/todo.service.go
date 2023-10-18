package todo

var todolist = []Todo{
	Todo{
		ID:     1,
		Task:   "Start",
		Status: true,
	},
	Todo{
		ID:     2,
		Task:   "Continue",
		Status: false,
	},
}

func List() []Todo {
	return todolist
}

func Add(todo Todo) []Todo {
	todolist = append(todolist, todo)
	return todolist
}
