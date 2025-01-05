// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: task.sql

package db

import (
	"context"
)

const completeTodo = `-- name: CompleteTodo :one
UPDATE todos SET  completed=$2 WHERE id = $1
RETURNING id, title, completed, created_at, updated_at
`

type CompleteTodoParams struct {
	ID        int32 `json:"id"`
	Completed bool  `json:"completed"`
}

func (q *Queries) CompleteTodo(ctx context.Context, arg CompleteTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, completeTodo, arg.ID, arg.Completed)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Completed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createTodo = `-- name: CreateTodo :one
INSERT INTO todos (title) VALUES ($1) RETURNING id, title, completed, created_at, updated_at
`

func (q *Queries) CreateTodo(ctx context.Context, title string) (Todo, error) {
	row := q.db.QueryRowContext(ctx, createTodo, title)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Completed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteTodoByID = `-- name: DeleteTodoByID :exec
DELETE  FROM todos WHERE id = $1
`

func (q *Queries) DeleteTodoByID(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTodoByID, id)
	return err
}

const getAllTodo = `-- name: GetAllTodo :many
SELECT id, title, completed, created_at, updated_at FROM todos
ORDER BY created_at DESC
LIMIT $1
OFFSET $2
`

type GetAllTodoParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetAllTodo(ctx context.Context, arg GetAllTodoParams) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, getAllTodo, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Todo{}
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Completed,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTodoByID = `-- name: GetTodoByID :one
SELECT id, title, completed, created_at, updated_at FROM todos
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetTodoByID(ctx context.Context, id int32) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getTodoByID, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Completed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateTodo = `-- name: UpdateTodo :one
UPDATE todos SET  title = $2 WHERE id = $1
RETURNING id, title, completed, created_at, updated_at
`

type UpdateTodoParams struct {
	ID    int32  `json:"id"`
	Title string `json:"title"`
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, updateTodo, arg.ID, arg.Title)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Completed,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
