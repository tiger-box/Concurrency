package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func process(ctx context.Context, wg *sync.WaitGroup) {
	result := make(chan int)
	defer wg.Done()
	go func() {
		time.Sleep(5 * time.Second)
		result <- 42
		close(result)
	}()

	select {

	case value, ok := <-result:
		if !ok {
			fmt.Println("closed")
			return
		}
		fmt.Println(value)
	case <-ctx.Done():
		fmt.Println("cancelled")
	}
}
func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	var wg sync.WaitGroup

	wg.Add(1)

	go process(ctx, &wg)

	wg.Wait()
	cancel()

}
