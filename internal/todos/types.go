package todos

type TodoDto struct {
	Title       string `json:"title" validate:""`
	Description string `json:"description" validate:""`
}

type Todo struct {
	ID          int    `json:"id" validate:""`
	Title       string `json:"title" validate:""`
	Description string `json:"description" validate:""`
}
