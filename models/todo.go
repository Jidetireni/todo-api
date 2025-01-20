package models

import (
	"time"

	"github.com/Jidetireni/todo-api/db"
)

type Todo struct {
	ID         int64      `json:"id"`
	Task       string     `json:"task"`
	Status     string     `json:"status"`
	CreatedAt  time.Time  `json:"created_at"`
	FinishedAT *time.Time `json:"finished_at"`
	UserID     int64      `json:"user_id"`
}

// binding:"required"

func (t *Todo) Save() error {

	if t.Status == "" {
		t.Status = "pending"
	}

	if t.CreatedAt.IsZero() {
		t.CreatedAt = time.Now()
	}

	query := `
	INSERT INTO todo (task, status, created_at, finished_at, user_id)
	VALUES (?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(t.Task, t.Status, t.CreatedAt, t.FinishedAT, t.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	t.ID = id
	if err != nil {
		return err
	}
	return nil
}

// func mapStatusToBool(status string) bool {
// 	return status == "completed"
// }

func GetAllTodolist() ([]Todo, error) {
	query := "SELECT * FROM todo"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var todolist []Todo

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Task, &todo.Status, &todo.CreatedAt, &todo.FinishedAT, &todo.UserID)
		if err != nil {
			return nil, err
		}
		todolist = append(todolist, todo)
	}

	return todolist, nil
}

func GetTodoById(id int64) (*Todo, error) {
	query := "SELECT * FROM todo WHERE id =?"
	row := db.DB.QueryRow(query, id)
	var todo Todo
	err := row.Scan(&todo.ID, &todo.Task, &todo.Status, &todo.CreatedAt, &todo.FinishedAT, &todo.UserID)
	if err != nil {
		return nil, err
	}

	return &todo, err
}

func (t Todo) Update() error {

	if t.Status == "" {
		t.Status = "pending"
	}

	if t.CreatedAt.IsZero() {
		t.CreatedAt = time.Now()
	}

	query := `
	UPDATE todo
	SET task = ?, status = ?, created_at = ? 
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(t.Task, t.Status, t.CreatedAt, t.ID)
	if err != nil {
		return err
	}
	return err
}

func (t Todo) Complete() error {

	var finishedAT time.Time
	if t.Status == "completed" {
		finishedAT = time.Now()
	}

	query := `
	UPDATE todo
	SET status = ?, finished_at = ? 
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.Status, finishedAT, t.ID)
	if err != nil {
		return err
	}

	return err
}

func (t Todo) Delete() error {
	query := "DELETE FROM todo WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(t.ID)
	return err
}
