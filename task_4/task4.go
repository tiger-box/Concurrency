package main

import (
	"fmt"
	"sync"
	"time"
)

const numWorkess = 3

func main() {
	task := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int)
	go func() {
		for i := 0; i < len(task); i++ {
			ch <- task[i]
		}
		close(ch)
	}()

	var wg sync.WaitGroup

	for i := 1; i <= (numWorkess); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range ch {
				fmt.Println("Worked", i, "processed task", job)
				time.Sleep(500 * time.Millisecond)
			}
		}()
	}

	wg.Wait()

}
