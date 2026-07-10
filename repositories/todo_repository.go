package repositories

import (
	"database/sql"

	"todolist-api/models"
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) FindAll() ([]models.Todo, error) {
	rows, err := r.db.Query("SELECT id, title, description, completed, created_at, updated_at FROM todos ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if todos == nil {
		todos = []models.Todo{}
	}

	return todos, nil
}

func (r *TodoRepository) FindByID(id int64) (*models.Todo, error) {
	var todo models.Todo
	err := r.db.QueryRow("SELECT id, title, description, completed, created_at, updated_at FROM todos WHERE id = ?", id).
		Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepository) Create(todo *models.Todo) error {
	result, err := r.db.Exec("INSERT INTO todos (title, description) VALUES (?, ?)", todo.Title, todo.Description)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	todo.ID = id
	return nil
}

func (r *TodoRepository) Update(todo *models.Todo) error {
	_, err := r.db.Exec("UPDATE todos SET title = ?, description = ?, completed = ? WHERE id = ?",
		todo.Title, todo.Description, todo.Completed, todo.ID)
	return err
}

func (r *TodoRepository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}

func (r *TodoRepository) MarkComplete(id int64) error {
	_, err := r.db.Exec("UPDATE todos SET completed = TRUE WHERE id = ?", id)
	return err
}

func (r *TodoRepository) MarkIncomplete(id int64) error {
	_, err := r.db.Exec("UPDATE todos SET completed = FALSE WHERE id = ?", id)
	return err
}