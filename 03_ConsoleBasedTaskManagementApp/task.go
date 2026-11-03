package main

import (
	"errors"
	"fmt"
)

type TaskStatus int

const (
	TaskStatusToDo = iota
	TaskStatusInProgress
	TaskStatusDone
)

func (s TaskStatus) String() string {
	switch s {
	case TaskStatusToDo:
		return "to do"
	case TaskStatusInProgress:
		return "in progress"
	case TaskStatusDone:
		return "done"
	default:
		return "unknown status"
	}
}

func (s TaskStatus) NextStatus() TaskStatus {
	switch s {
	case TaskStatusToDo:
		return TaskStatusInProgress
	case TaskStatusInProgress:
		return TaskStatusDone
	default:
		return s
	}
}

func TaskStatusFromString(statusString string) (TaskStatus, error) {
	switch statusString {
	case "todo":
		return TaskStatusToDo, nil
	case "inprogress":
		return TaskStatusInProgress, nil
	case "done":
		return TaskStatusDone, nil
	default:
		return TaskStatusToDo, errors.New("unknown status")
	}
}

type Task struct {
	ID     int
	Title  string
	Status TaskStatus
}

func NewTask(title string) *Task {
	return &Task{
		ID:     -1,
		Title:  title,
		Status: TaskStatusToDo,
	}
}

func (t *Task) String() string {
	return fmt.Sprintf("#%d %s [%s]", t.ID, t.Title, t.Status)
}
