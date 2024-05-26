package data

import "database/sql"

type DAO struct {
	Todos *TodoDAO
}

func NewDAO(db *sql.DB) *DAO {
	return &DAO{
		Todos: &TodoDAO{DB: db},
	}
}
