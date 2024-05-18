package main

import (
	"fmt"
	"time"
)

func main() {
	waitTime := 1 * time.Second
	for {
		fmt.Printf("Hello, wait %v\n", waitTime)
		time.Sleep(waitTime)
		waitTime *= 2 // increase exponentially
	}
}
