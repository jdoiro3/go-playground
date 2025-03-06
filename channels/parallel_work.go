package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func defaultWorker(i int) int {
	sleepFor := time.Duration(rand.Intn(10)) * time.Second
	time.Sleep(sleepFor)
	fmt.Printf("worker slept for %v...\n", sleepFor)
	return i * i
}

func parallelWork[T any, R any](data []T, worker func(T) R) <-chan R {
	results := make(chan R)
	var wg sync.WaitGroup
	for _, i := range data {
		wg.Add(1)
		go func(i T) {
			defer wg.Done()
			results <- worker(i)
		}(i)
	}
	go func(wg *sync.WaitGroup, results chan R) {
		wg.Wait()
		close(results)
	}(&wg, results)
	return results
}

func main() {
	for r := range parallelWork[int, int](makeRange(1, 20), defaultWorker) {
		fmt.Println(r)
	}
}
