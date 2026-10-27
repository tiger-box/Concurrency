package main

import (
	"context"
	"fmt"
	"sync"
)

type Task struct {
	id     int
	result int
}

func ProcessBatch(ctx context.Context, values []int) ([]int, error) {
	var wg sync.WaitGroup

	ch := make(chan Task, len(values))

	for i, v := range values {
		wg.Add(1)
		go func(id int, value int) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			case ch <- Task{id: id, result: value * 2}:
				return
			}
		}(i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	result := make([]int, len(values))

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case c, ok := <-ch:
			if !ok {
				return result, nil
			}
			result[c.id] = c.result
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	values := []int{1, 2, 3, 4, 5}

	result, err := ProcessBatch(ctx, values)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Result: %v\n", result)
}
