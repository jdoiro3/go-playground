package main

import (
	"math/rand"
	"sync"
	"time"
)

func randSleep(max int) int {
	r := rand.Intn(max)
	time.Sleep(time.Duration(r) * time.Second)
	return r
}

func sendNums(nums chan<- int, total int, maxSleep int) {
	var wg sync.WaitGroup
	wg.Add(total + 1) // the loop will run total + 1 times
	for i := 0; i <= total; i++ {
		go func(i int) {
			defer wg.Done()
			sleptFor := randSleep(maxSleep)
			nums <- i
			print("Slept for ", sleptFor, " seconds.\n")
		}(i)
	}

	go func(wg *sync.WaitGroup) {
		wg.Wait()
		print("Done sending nums. Closing channel\n")
		close(nums)
	}(&wg)
}

func main() {
	var wg sync.WaitGroup
	nums := make(chan int)

	go sendNums(nums, 10, 5)

	wg.Add(1)
	go func(nums <-chan int) {
		defer wg.Done()
		for x := range nums {
			print(x, "\n")
		}
	}(nums)
	wg.Wait()
}
