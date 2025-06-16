package main

import (
	"context"
	"fmt"
	"time"
)

func generator(ctx context.Context, numb ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range numb {
			select {
			case <-ctx.Done():
				fmt.Println("Entered Early")
				return
			case out <- n:
			}
		}
	}()
	return out
}
func square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for {

			select {
			case <-ctx.Done():
				fmt.Println("Exited Early")
				return
			case n, ok := <-in:
				if !ok {
					return
				}
				out <- n * n
			}
		}
	}()
	return out
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	in := generator(ctx, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	out := square(ctx, in)

	for n := range out {
		time.Sleep(600 * time.Millisecond)
		fmt.Println(n)
	}
}
