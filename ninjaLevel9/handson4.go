package main

import (
	"fmt"
	"sync"
)

func main() {
	//fix the race condition using mutex

	var ws sync.WaitGroup

	inc := 0
	gs := 100

	ws.Add(gs)

	var m sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			m.Lock()
			v := inc
			v++
			inc = v
			fmt.Println("on thread:", inc)
			m.Unlock()

			ws.Done()
		}()
	}

	ws.Wait()
	fmt.Println("last inc val is,", inc)
	fmt.Println("main Done")

}
