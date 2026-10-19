package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Worker(ctx context.Context, jobs <-chan int, wg *sync.WaitGroup, count int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker stopped:", count)
			return
			// Подсказал GPT - нужно обьяснение
			//Почему нужен тут еще select???
		default:
		}

		select {
		case <-ctx.Done():
			fmt.Println("worker stopped:")
			return

		case j, ok := <-jobs:
			if !ok {
				fmt.Println("worker stoped:")
				return
			}
			fmt.Println("worker:", count, "job:", j)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {

	const numberJob = 100
	ctx, cancel := context.WithCancel(context.Background())
	jobs := make(chan int, numberJob)

	go func() {
		task := 1
		for i := 0; i < numberJob; i++ {
			jobs <- task
			task++
		}
		close(jobs)
	}()

	var wg sync.WaitGroup

	workers := 3
	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go Worker(ctx, jobs, &wg, i)
	}

	time.Sleep(5 * time.Second)

	cancel()

	wg.Wait()

}
