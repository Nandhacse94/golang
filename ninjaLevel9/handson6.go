package main

import (
	"fmt"
	"runtime"
)

func main() {
	//print OS and ARCH
	fmt.Println("Current OS is, ",runtime.GOOS)
	fmt.Println("Current ARCH is, ",runtime.GOARCH)
}

