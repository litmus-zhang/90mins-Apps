package db

import (
	"context"
	"testing"

	"github.com/litmus-zhang/90min-app-todo/util"
	"github.com/stretchr/testify/require"
)

func TestCreateTodo(t *testing.T) {

	tests := []struct {
		name string
		todo string
	}{
		{
			name: "test normal todo",
			todo: util.RandomString(15),
		},
		{
			name: "test empty todo",
			todo: "",
		},
		{
			name: "test todo with special characters",
			todo: "test@#$%^&*()",
		},
		{
			name: "test todo with numbers",
			todo: "1234567890",
		},
		{
			name: "test todo with spaces",
			todo: "test   with  	spaces",
		},
		{
			name: "test todo with new lines",
			todo: "test\nwith\nnew\nlines",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			todo, err := testQueries.CreateTodo(context.Background(), tt.todo)
			require.NoError(t, err)
			require.NotEmpty(t, todo)
			require.Equal(t, tt.todo, todo.Title)
		})
	}

}

func TestGetTodoByID(t *testing.T) {
	todo, err := testQueries.CreateTodo(context.Background(), util.RandomString(15))
	require.NoError(t, err)
	require.NotEmpty(t, todo)

	todoByID, err := testQueries.GetTodoByID(context.Background(), todo.ID)
	require.NoError(t, err)
	require.NotEmpty(t, todoByID)
	require.Equal(t, todo.ID, todoByID.ID)
	require.Equal(t, todo.Title, todoByID.Title)
}
func TestListTodos(t *testing.T) {
	for i := 0; i < 10; i++ {
		_, err := testQueries.CreateTodo(context.Background(), util.RandomString(15))
		require.NoError(t, err)
	}
	args := GetAllTodoParams{
		Limit:  10,
		Offset: 0,
	}
	todos, err := testQueries.GetAllTodo(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, todos, 10)
}

func TestUpdateTodoStatus(t *testing.T) {
	todo, err := testQueries.CreateTodo(context.Background(), util.RandomString(15))
	require.NoError(t, err)
	require.NotEmpty(t, todo)

	args := UpdateTodoParams{
		ID:    todo.ID,
		Title: util.RandomString(15) + " updated",
	}

	todo2, err := testQueries.UpdateTodo(context.Background(), args)
	require.NoError(t, err)

	require.NotEmpty(t, todo2)
	require.Equal(t, args.Title, todo2.Title)
	require.Equal(t, todo.ID, todo2.ID)

}

func TestCompleteTodo(t *testing.T) {
	todo, err := testQueries.CreateTodo(context.Background(), util.RandomString(15))
	require.NoError(t, err)
	require.NotEmpty(t, todo)
	args := CompleteTodoParams{
		ID:        todo.ID,
		Completed: true,
	}

	todo2, err := testQueries.CompleteTodo(context.Background(), args)
	require.NoError(t, err)

	require.NotEmpty(t, todo2)
	require.Equal(t, todo.ID, todo2.ID)
	require.True(t, todo2.Completed)
}

func TestDeleteTodoByID(t *testing.T) {
	todo, err := testQueries.CreateTodo(context.Background(), util.RandomString(15))
	require.NoError(t, err)
	require.NotEmpty(t, todo)

	err = testQueries.DeleteTodoByID(context.Background(), todo.ID)
	require.NoError(t, err)

	todoByID, err := testQueries.GetTodoByID(context.Background(), todo.ID)
	require.Error(t, err)
	require.Empty(t, todoByID)
}
