package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Process(numbers []int) []int {
	var wg sync.WaitGroup
	ch := make(chan int)
	arrayCh := []int{}

	if len(numbers) == 0 {
		return nil
	}

	for i := range numbers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			times := time.Duration(rand.Intn(401)+100) * time.Millisecond
			numbers[i] = numbers[i] * 2
			time.Sleep(times)
			ch <- numbers[i]
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	for value := range ch {
		arrayCh = append(arrayCh, value)
	}

	return arrayCh
}

func main() {
	var array []int = []int{1, 2, 3, 4}
	fmt.Println(Process(array))

}
