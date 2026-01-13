package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Go(func() {
		defer func() {
			_ = recover()
		}()

		c := make(chan struct{}, 10)
		close(c)

		select {
		case c <- struct{}{}:
			fmt.Println("sent to channel")

		default:
		}

		fmt.Println("done")
	})

	wg.Wait()
	fmt.Println("done")
}
