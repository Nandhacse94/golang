package main

import (
	"fmt"
	"sync/atomic"
	"sync"
)

func main() {
	//fix the race condition using Atomic

	var ws sync.WaitGroup

	var inc int64
	gs := 100

	ws.Add(gs)

	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&inc,1)
			fmt.Println("on thread:", atomic.LoadInt64(&inc))
			
			ws.Done()
		}()
	}

	ws.Wait()
	fmt.Println("last inc val is,", inc)
	fmt.Println("main Done")

}

