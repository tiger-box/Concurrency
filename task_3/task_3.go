package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Worker(ctx context.Context, wg *sync.WaitGroup) {
	ticker := time.NewTicker(1000 * time.Millisecond)
	defer ticker.Stop()
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker stopped")
			return

		case <-ticker.C:
			fmt.Println("worker is working")
		}
	}

}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go Worker(ctx, &wg)

	time.Sleep(5 * time.Second)

	cancel()
	wg.Wait()

}
