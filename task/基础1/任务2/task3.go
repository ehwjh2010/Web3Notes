package main

import (
	"fmt"
	"sync"
	"time"
)

func GenerateNum(start, end int, ch chan<- int) {
	for i := start; i <= end; i++ {
		ch <- i
	}
	close(ch)
}

func PrintNum(ch <-chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}

func DoNumTask(start, end int, ch chan int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		GenerateNum(start, end, ch)
		wg.Done()
	}()

	go func() {
		PrintNum(ch)
		wg.Done()
	}()

	wg.Wait()
}

func main() {
	DoNumTask(1, 10, make(chan int))

	time.Sleep(10 * time.Second)

	DoNumTask(1, 100, make(chan int, 30))
}
