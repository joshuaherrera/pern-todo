package psql

import (
	"database/sql"

	"github.com/joshuaherrera/pern-todo/go_server/go_server/models"
)

// TodoModel defines wrapper for sql.DB connection pool
type TodoModel struct {
	DB *sql.DB
}

// All grabs all todos from db
func (m *TodoModel) All() ([]*models.Todo, error) {
	stmt := `SELECT todo_id, description FROM todo`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todos := []*models.Todo{}

	for rows.Next() {
		t := &models.Todo{}

		err := rows.Scan(&t.TodoID, &t.Description)
			if err != nil {
				return nil, err
		}

		todos = append(todos, t)
	}
	// check for errors in scanning
		if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}