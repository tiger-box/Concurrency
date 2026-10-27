package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Result struct {
	Value int
	Err   error
}

func Worker(ctx context.Context, jobs <-chan int, result chan<- Result, wg *sync.WaitGroup, id int) {

	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker canceled:")
			return

		case j, ok := <-jobs:
			if !ok {
				fmt.Println("worker stopped")
				return
			}
			if j%2 == 0 {
				result <- Result{Value: j * 2, Err: nil}
				fmt.Println("worker", id, "job", j)
			} else {
				result <- Result{Err: fmt.Errorf("odd number:")}
			}
		}
	}
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	const numbersJobs = 20
	jobs := make(chan int, numbersJobs)
	result := make(chan Result, numbersJobs)

	task := 0
	for i := 1; i <= numbersJobs; i++ {
		task++
		jobs <- task
	}

	close(jobs)
	defer cancel()

	var wg sync.WaitGroup
	for i := 1; i <= numbersJobs; i++ {
		wg.Add(1)
		go Worker(ctx, jobs, result, &wg, i)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	for res := range result {
		if res.Err != nil {
			fmt.Println("Error:", res.Err)
		} else {
			fmt.Println("Result:", res.Value)
		}
	}
}
