package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//waitgroup,bunch of goroutine to see race condition.

	var ws sync.WaitGroup

	inc := 0
	gs := 100

	ws.Add(gs)

	for i := 0; i < 100; i++ {
		go func() {
			v := inc
			runtime.Gosched()
			v++
			inc = v
			fmt.Println("on thread:", inc)
			ws.Done()
		}()
	}

	ws.Wait()
	fmt.Println("last inc val is,", inc)
	fmt.Println("main Done")

}
