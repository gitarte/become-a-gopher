package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	globalChan := make(chan int, 10)
	globalTimeOut := make(chan struct{})

	wg.Add(1)
	go func(ch chan<- struct{}) {
		defer wg.Done()

		time.Sleep(3 * time.Second)

		ch <- struct{}{}
	}(globalTimeOut)

	//reader
	wg.Add(1)
	go func(ch <-chan int, timeout <-chan struct{}) {
		defer wg.Done()
		for {
			select {
			case i, ok := <-ch:
				if ok {
					fmt.Println(i)
				}
			case <-timeout:
				return
			}
		}
	}(globalChan, globalTimeOut)

	//writer
	wg.Add(1)
	go func(ch chan<- int) {
		defer wg.Done()
		ch <- 123
		ch <- 124
		close(ch)
	}(globalChan)

	wg.Wait()
	fmt.Println("koniec")
}
