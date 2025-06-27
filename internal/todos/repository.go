package todos

import (
	"errors"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/umdalecs/todo-list-api/internal/db"
)

var (
	ErrTodoNotFound = errors.New("the specified todo does not exist")
)

type TodoRepository struct {
	db *pgxpool.Pool
}

func NewTodoRepository(db *pgxpool.Pool) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (r *TodoRepository) CreateTodo(userID int, dto TodoDto, todo *Todo) error {
	err := pgxscan.Get(db.Ctx, r.db, todo, `
		INSERT INTO todos(title, description, user_id) 
		VALUES ($1, $2, $3) 
		RETURNING id, title, description`,
		dto.Title, dto.Description, userID)

	if err != nil {
		// I should manage the user not exists edge
		// case but i will not because the userID
		// value comes from a valid jwt
		return err
	}

	return nil
}

func (r *TodoRepository) UpdateTodo(todoID, userID int, dto TodoDto, todo *Todo) error {
	err := pgxscan.Get(db.Ctx, r.db, todo, `
		UPDATE todos 
		SET title = $1, description = $2 
		WHERE id = $3 AND user_id = $4 
		RETURNING id, title, description`,
		dto.Title, dto.Description, todoID, userID)

	if err != nil {
		return err
	}

	if todo.ID == 0 {
		return ErrTodoNotFound
	}

	return nil
}

func (r *TodoRepository) DeleteTodo(userID, todoID int) error {
	var todo Todo

	err := pgxscan.Get(db.Ctx, r.db, &todo, `
		SELECT id FROM todos 
		WHERE id = $1 AND user_id = $2`,
		userID, todoID)

	if err != nil {
		return err
	}

	if todo.ID == 0 {
		return ErrTodoNotFound
	}

	_, err = r.db.Exec(db.Ctx, `
		DELETE FROM todos 
		WHERE id = $1`, todoID)

	if err != nil {
		return err
	}

	return nil
}

func (r *TodoRepository) GetAllTodos(userID, page, limit int, filter string, todos *[]Todo, totalPages *int) error {
	offset := limit * (page - 1)

	var total int
	rows, err := r.db.Query(db.Ctx, `
		SELECT count(*)
		FROM todos
		WHERE user_id = $1
			AND ( title ILIKE '%' || $2 || '%'
			OR description ILIKE '%' || $2 || '%' )
	`, userID, filter)

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			return err
		}
	}

	*totalPages = total / limit

	if total%limit != 0 {
		*totalPages++
	}

	err = pgxscan.Select(db.Ctx, r.db, todos, `
		SELECT id, title, description
		FROM todos
		WHERE user_id = $1
			AND ( title ILIKE '%' || $4 || '%'
			OR description ILIKE '%' || $4 || '%' )
		ORDER BY id
		LIMIT $2
		OFFSET $3
	`, userID, limit, offset, filter)

	if err != nil {
		return err
	}

	return nil
}
