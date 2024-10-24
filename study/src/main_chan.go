package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg = sync.WaitGroup{}
	ch := make(chan int, 100)
	wg.Add(2)
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 20
		ch <- 21
		close(ch) // thong bao khi khong con gui nua
		wg.Done()
	}(ch)
	wg.Wait()
}
