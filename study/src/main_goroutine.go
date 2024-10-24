package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg = sync.WaitGroup{}
	wg.Add(1)
	go func() {
		count(19)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Study")
}

func count(number int) {
	for i := 0; i <= number; i++ {
		fmt.Println("Count", i)
		time.Sleep(time.Second)
	}
}
