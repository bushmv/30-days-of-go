package main

import (
	"fmt"
	"time"
)

type TaskResult[R any] struct {
	TaskId   int
	WorkerId int
	Value    R
	Err      error
	Elapsed  time.Duration
}

func (r *TaskResult[R]) String() string {
	if r.Err != nil {
		return fmt.Sprintf("task with id %d failed, err: %v (took time: %v, workerId: %d)", r.TaskId, r.Err, r.Elapsed, r.WorkerId)
	}
	return fmt.Sprintf("task with id %d completed successfully with result %v (took time: %v, workerId: %d)", r.TaskId, r.Value, r.Elapsed, r.WorkerId)
}
