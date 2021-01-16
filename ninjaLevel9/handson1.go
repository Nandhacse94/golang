package main

import (
	"fmt"
	"runtime"
	"sync"
)

//use waitgroups to make sure each goroutine finished before your program exists.
var ws sync.WaitGroup

func foo() {
	fmt.Println("Executing Foo!")
	ws.Done()
}

func bar() {
	fmt.Println("Executing Bar!")
	ws.Done()
}

func main() {
	fmt.Println("Num CPU: ", runtime.NumCPU())
	fmt.Println("Num go routine: ", runtime.NumGoroutine())

	//In addition to main goroutine, call two more goroutine.
	ws.Add(2)

	go foo()
	go bar()

	fmt.Println("Num go routine: ", runtime.NumGoroutine())
	ws.Wait()

	fmt.Println("function main done")
}

