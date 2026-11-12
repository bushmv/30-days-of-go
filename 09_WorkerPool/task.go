package main

import "context"

type Task[T any] struct {
	ID   int
	ctx  context.Context
	Data T
}

func NewTask[T any](ctx context.Context, id int, data T) *Task[T] {
	return &Task[T]{
		ID:   id,
		ctx:  ctx,
		Data: data,
	}
}
