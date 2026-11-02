package main

import (
	"errors"
)

var ErrCircularBufferIsEmpty = errors.New("circular buffer is empty")
var ErrCircularBufferIsFull = errors.New("circular buffer is full")

type CircularBuffer[T any] struct {
	buf  []T
	len  int
	cap  int
	head int
	tail int
}

func NewCircularBuffer[T any](cap int) *CircularBuffer[T] {
	return &CircularBuffer[T]{
		len:  0,
		cap:  cap,
		buf:  make([]T, cap),
		head: 0,
		tail: 0,
	}
}

func (c *CircularBuffer[T]) Put(v T) error {
	if c.IsFull() {
		return ErrCircularBufferIsFull
	}
	c.buf[c.head] = v
	c.head = (c.head + 1) % c.cap
	c.len++
	return nil
}

func (c *CircularBuffer[T]) Get() (T, error) {
	if c.IsEmpty() {
		var zero T
		return zero, ErrCircularBufferIsEmpty
	}
	value := c.buf[c.tail]
	c.tail = (c.tail + 1) % c.cap
	c.len--
	return value, nil
}

func (c *CircularBuffer[T]) Len() int {
	return c.len
}

func (c *CircularBuffer[T]) IsEmpty() bool {
	return c.len == 0
}

func (c *CircularBuffer[T]) IsFull() bool {
	return c.len == c.cap
}

func (c *CircularBuffer[T]) Dump() []T {
	values := make([]T, c.cap)
	copy(values, c.buf)
	return values
}

func (c *CircularBuffer[T]) Values() []T {
	values := make([]T, 0, c.len)
	for i := range c.len {
		idx := (i + c.tail) % c.cap
		values = append(values, c.buf[idx])
	}
	return values
}
