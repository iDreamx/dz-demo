package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	numGoroutines := 3
	code := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		ot := 0 + (i * 4)
		do := 4 + (i * 4)
		go func() {
			getSum(code, arr[ot:do])
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(code)
	}()

	sum := 0
	for res := range code {
		sum += res
	}
	fmt.Println("Sum:", sum)
}

func getSum(codeCh chan int, arr []int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	codeCh <- sum
}
