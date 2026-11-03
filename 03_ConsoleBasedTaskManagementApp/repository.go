package main

import (
	"database/sql"
	"fmt"
)

const createTasksQuery = `CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    status INTEGER DEFAULT 0
);`
const insertTaskQuery = `INSERT INTO tasks (title, status) VALUES (?, ?);`
const selectTaskByIdQuery = `SELECT id, title, status FROM tasks WHERE id = ?;`
const selectAllTasksQuery = `SELECT id, title, status FROM tasks;`
const selectTasksWithStatusQuery = `SELECT id, title, status FROM tasks WHERE status = ?;`
const updateTaskStatusQuery = `UPDATE tasks SET status = ? WHERE id = ?;`
const deleteTaskByIdQuery = `DELETE FROM tasks WHERE id = ?;`

type Repository struct {
	db *sql.DB
}

func NewRepository(dbpath string) (*Repository, error) {
	db, err := sql.Open("sqlite3", dbpath)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(createTasksQuery)
	if err != nil {
		return nil, err
	}
	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) Close() error {
	return r.db.Close()
}

func (r *Repository) CreateTask(title string) (*Task, error) {
	task := NewTask(title)
	res, err := r.db.Exec(insertTaskQuery, task.Title, task.Status)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	task.ID = int(id)
	return task, nil
}

func (r *Repository) GetTaskById(taskId int) (*Task, error) {
	var task Task
	if err := r.db.QueryRow(selectTaskByIdQuery, taskId).Scan(&task.ID, &task.Title, &task.Status); err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *Repository) GetAllTasks() ([]*Task, error) {
	rows, err := r.db.Query(selectAllTasksQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := make([]*Task, 0)
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *Repository) GetTasksByStatus(status TaskStatus) ([]*Task, error) {
	rows, err := r.db.Query(selectTasksWithStatusQuery, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := make([]*Task, 0)
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *Repository) ChangeTaskStatus(taskId int, newStatus TaskStatus) error {
	res, err := r.db.Exec(updateTaskStatusQuery, newStatus, taskId)
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return fmt.Errorf("task #%d not found", taskId)
	}
	return nil
}

func (r *Repository) DeleteTaskById(taskId int) error {
	res, err := r.db.Exec(deleteTaskByIdQuery, taskId)
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return fmt.Errorf("task #%d not found", taskId)
	}
	return nil
}
