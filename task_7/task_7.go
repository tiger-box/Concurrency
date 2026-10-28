package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func process(ctx context.Context) {
	result := make(chan int)
	defer close(result)

	go func() {
		select {
		case <-time.After(5 * time.Second):
			result <- 42

		case <-ctx.Done():
			fmt.Println("can")
			return
		}
	}()

	select {
	case value := <-result:
		fmt.Println(value)
	case <-ctx.Done():
		fmt.Println("cancelled")
	}

}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	process(ctx)

	time.Sleep(2 * time.Second)
	fmt.Println(runtime.NumGoroutine())

}
