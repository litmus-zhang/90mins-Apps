-- name: CreateTodo :one
INSERT INTO todos (title) VALUES ($1) RETURNING *;

-- name: GetTodoByID :one
SELECT * FROM todos
WHERE id = $1
LIMIT 1;

-- name: UpdateTodo :one
UPDATE todos SET  title = $2 WHERE id = $1
RETURNING *;

-- name: CompleteTodo :one
UPDATE todos SET  completed=$2 WHERE id = $1
RETURNING *;

-- name: GetAllTodo :many
SELECT * FROM todos
ORDER BY created_at DESC
LIMIT $1
OFFSET $2;


-- name: DeleteTodoByID :exec
DELETE  FROM todos WHERE id = $1;

