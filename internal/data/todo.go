package data

import (
	"database/sql"
	_ "embed"
	"time"
)

//go:embed queries/insert_todo.sql
var insertTodoQuery string

//go:embed queries/get_todo.sql
var getTodoQuery string

//go:embed queries/complete_todo.sql
var completeTodoQuery string

//go:embed queries/delete_todo.sql
var deleteTodoQuery string

//go:embed queries/get_todos.sql
var getTodosQuery string

type Todo struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Content     string     `json:"content"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

type TodoDAO struct {
	DB *sql.DB
}

func (dao *TodoDAO) Insert(todo *Todo) error {
	args := []any{todo.Title, todo.Content}
	err := dao.DB.QueryRow(insertTodoQuery, args...).Scan(&todo.ID, &todo.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (dao *TodoDAO) Get(id int64) (*Todo, error) {
	var todo Todo
	err := dao.DB.QueryRow(getTodoQuery, id).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Content,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.CompletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (dao *TodoDAO) GetMultiple(limit int, offset int) ([]*Todo, error) {
	todos := []*Todo{}
	rows, err := dao.DB.Query(getTodosQuery, limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Content,
			&todo.Completed,
			&todo.CreatedAt,
			&todo.CompletedAt,
		)
		if err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}
	return todos, nil
}

func (dao *TodoDAO) Complete(id int64) error {
	_, err := dao.DB.Exec(completeTodoQuery, id)
	if err != nil {
		return err
	}
	return nil
}

func (dao *TodoDAO) Delete(id int64) error {
	_, err := dao.DB.Exec(deleteTodoQuery, id)
	if err != nil {
		return err
	}
	return nil
}
