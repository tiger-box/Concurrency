package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 10
		ch <- 20
		ch <- 30
		// close(ch) - можно и так
	}()

	//ждет паралельно пока все горутины не закончат работу.
	go func() {
		wg.Wait()
		close(ch)
	}()

	for value := range ch {
		fmt.Println(value)
	}

	// wg.Wait

}
