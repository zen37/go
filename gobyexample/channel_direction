package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i < 40; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 40; j++ {
		fmt.Println(j, " is ", <-results)
	}

	fmt.Println("Milliseconds: ", time.Since(start).Milliseconds())
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
