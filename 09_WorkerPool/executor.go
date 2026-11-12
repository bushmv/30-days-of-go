package main

import (
	"errors"
	"sync"
	"time"
)

type Executor[T any, R any] struct {
	running    bool
	mu         sync.Mutex
	handleFunc func(task *Task[T]) (R, error)
	tasks      chan *Task[T]
	results    chan *TaskResult[R]
	stopCh     chan struct{}
	wg         sync.WaitGroup
}

func NewExecutor[T any, R any](handleFunc func(task *Task[T]) (R, error)) *Executor[T, R] {
	return &Executor[T, R]{
		handleFunc: handleFunc,
	}
}

func (e *Executor[T, R]) run(workerId int) {
	defer e.wg.Done()
	for task := range e.tasks {
		start := time.Now()
		value, err := e.handleFunc(task)
		elapsed := time.Since(start)
		result := &TaskResult[R]{
			TaskId:   task.ID,
			WorkerId: workerId,
			Elapsed:  elapsed,
		}
		if err != nil {
			result.Err = err
		} else if task.ctx.Err() != nil {
			result.Err = task.ctx.Err()
		} else {
			result.Value = value
		}
		e.results <- result
	}
}

func (e *Executor[T, R]) Submit(task *Task[T]) error {
	e.mu.Lock()
	if !e.running {
		e.mu.Unlock()
		return errors.New("executor not running")
	}
	e.mu.Unlock()
	select {
	case e.tasks <- task:
		return nil
	case <-e.stopCh:
		return errors.New("executor stopped")
	}
}

func (e *Executor[T, R]) Start(workersCount int) (<-chan *TaskResult[R], error) {
	e.mu.Lock()
	if e.running {
		return nil, errors.New("executor already running")
	}
	defer e.mu.Unlock()
	e.running = true
	e.tasks = make(chan *Task[T], workersCount)
	e.results = make(chan *TaskResult[R])
	e.stopCh = make(chan struct{})
	e.wg.Add(workersCount)
	for workerId := range workersCount {
		go e.run(workerId)
	}
	return e.results, nil
}

func (e *Executor[T, R]) Stop() {
	e.mu.Lock()
	if !e.running {
		e.mu.Unlock()
		return
	}
	e.running = false
	e.mu.Unlock()
	close(e.stopCh)
	close(e.tasks)
	e.wg.Wait()
	close(e.results)
}
