package todos

import (
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/umdalecs/todo-list-api/internal/db"
)

type TodoRepository struct {
	db *pgxpool.Pool
}

func NewTodoRepository(db *pgxpool.Pool) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) CreateTodo(id int, dto TodoDto, todo *Todo) error {
	err := pgxscan.Get(db.Ctx, r.db, todo, `INSERT INTO todos(title, description, user_id) VALUES ($1, $2, $3) RETURNING *`, dto.Title, dto.Description, id)

	if err != nil {
		// I should manage the user not exists edge case
		// But i will not because it comes from a valid jwt
		return err
	}

	return nil
}

func (r *TodoRepository) UpdateTodo(id, userID int, dto TodoDto, todo *Todo) error {
	err := pgxscan.Get(db.Ctx, r.db, todo, `UPDATE todos SET title = $1, description = $2 WHERE id = $3 AND user_id = $4 RETURNING *`, dto.Title, dto.Description, id, userID)

	if err != nil {
		// TODO: Maybe an edge case is "todo doesn't exist"
		// and i should custom handle it
		return err
	}

	return nil
}

func (r *TodoRepository) DeleteTodo(id int) error {
	return nil
}

func (r *TodoRepository) GetAllTodos(id int, filter string, todos *[]Todo) error {
	return nil
}
