package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

func sendTasks(N int, e *Executor[int, int]) {
	go func() {
		for i := range N {
			ctx := context.Background()
			task := NewTask(ctx, i, i)
			e.Submit(task)
		}
	}()
}

func handleFunc(task *Task[int]) (int, error) {
	time.Sleep(time.Second)
	if rand.Float64() < 0.01 {
		return 0, fmt.Errorf("task with id %d failed", task.ID)
	}
	return task.Data, nil
}

func main() {
	e := NewExecutor(handleFunc)
	results, err := e.Start(2)
	if err != nil {
		panic(err)
	}
	sendTasks(5, e)
	for result := range results {
		fmt.Println(result)
	}
}
