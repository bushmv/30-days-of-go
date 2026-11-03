package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

const version = "1.0.0"

const help = `tasks usage:

    - tasks add "task title" - add new task with title "task title"
	- tasks id 42 - show task with id 42
    - tasks list - show all tasks with their status
    - tasks status [todo|inprogress|done] - show only tasks with specified status
    - tasks up 4 - update status of task with id to the next status (to do -> in progress -> done). Task with status "done" status not change
    - tasks delete "4" - delete task with id 4
    - tasks help - this message
    - tasks version - version of this app`

func main() {
	repository, err := NewRepository("data.db")
	if err != nil {
		fmt.Println("cannot init sqlite database")
		os.Exit(1)
	}
	defer repository.Close()
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("command not specified")
		fmt.Println(help)
		os.Exit(1)
	}
	cmd := args[0]
	switch cmd {
	case "add":
		if len(args) < 2 {
			fmt.Println("title for new task not specified")
			os.Exit(1)
		}
		title := args[1]
		task, err := repository.CreateTask(title)
		if err != nil {
			fmt.Printf("create task error: %s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("task \"%s\" created with id #%d\n", task.Title, task.ID)
	case "id":
		taskId, err := parseTaskId(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		task, err := repository.GetTaskById(taskId)
		if err != nil {
			fmt.Printf("load task error: %s\n", err.Error())
			os.Exit(1)
		}
		fmt.Printf("#%d %s [%s]\n", task.ID, task.Title, task.Status.String())
	case "list":
		tasks, err := repository.GetAllTasks()
		if err != nil {
			fmt.Printf("load task error: %s", err.Error())
			os.Exit(1)
		}
		n := len(tasks)
		fmt.Println("--- all tasks ---")
		for i, task := range tasks {
			fmt.Printf("[%d/%d] #%d %s [%s]\n", i+1, n, task.ID, task.Title, task.Status.String())
		}
	case "status":
		if len(args) < 2 {
			fmt.Println("status not specified")
			os.Exit(1)
		}
		statusString := args[1]
		status, err := TaskStatusFromString(statusString)
		if err != nil {
			fmt.Printf("unknown status: %q\n", statusString)
			os.Exit(1)
		}
		tasks, err := repository.GetTasksByStatus(status)
		if err != nil {
			fmt.Printf("load tasks error: %s", err.Error())
			os.Exit(1)
		}
		n := len(tasks)
		fmt.Printf("--- tasks in status %s ---\n", statusString)
		for i, task := range tasks {
			fmt.Printf("[%d/%d] #%d %s\n", i+1, n, task.ID, task.Title)
		}
	case "up":
		// no race condition between get task and update task because it's CLI tool
		taskId, err := parseTaskId(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		task, err := repository.GetTaskById(taskId)
		if err != nil {
			fmt.Printf("load task error: %s", err.Error())
			os.Exit(1)
		}
		if task.Status == TaskStatusDone {
			fmt.Printf("task #%d already in status done\n", taskId)
			return
		}
		newStatus := task.Status.NextStatus()
		if err := repository.ChangeTaskStatus(taskId, newStatus); err != nil {
			fmt.Printf("update status error: %s", err.Error())
			os.Exit(1)
		}
		fmt.Printf("status of task #%d updated to %s\n", task.ID, newStatus.String())
	case "delete":
		taskId, err := parseTaskId(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := repository.DeleteTaskById(taskId); err != nil {
			fmt.Printf("delete task error: %s", err.Error())
			os.Exit(1)
		}
		fmt.Printf("task #%d deleted\n", taskId)
	case "version":
		fmt.Printf("tasks version: %s\n", version)
	case "help":
		fmt.Println(help)
	default:
		fmt.Printf("unknown command: %q\n", cmd)
		fmt.Println(help)
		os.Exit(1)
	}
}

func parseTaskId(args []string) (int, error) {
	if len(args) < 2 {
		return 0, errors.New("task id not specified")
	}
	taskId, err := strconv.Atoi(args[1])
	if err != nil {
		return 0, errors.New("task id not integer")
	}
	return taskId, nil
}
