package main

import (
	"fmt"
	"sync"
)

func generator(numb ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range numb {
			out <- n
		}
		close(out)
	}()
	return out
}
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func squareWorkerPool(in <-chan int, numberOfWorkers int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	worker := func() {
		defer wg.Done()
		for n := range in {
			out <- n * n
		}
	}

	for i := 0; i < numberOfWorkers; i++ {
		wg.Add(1)
		go worker()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	in := generator(1, 2, 3, 4, 5)
	in2 := generator(10, 20, 30, 40, 50)
	out := square(in)
	out2 := squareWorkerPool(in2, 4)
	for n := range out {
		fmt.Println(n)
	}
	fmt.Println("--=--")
	for n := range out2 {
		fmt.Println(n)
	}

}
