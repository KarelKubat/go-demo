package main

import (
	"time"
	"fmt"
)

func longRunningCall(doneChannel chan bool) {
	fmt.Printf("starting long-ish call\n")
	time.Sleep(time.Second)
	fmt.Printf("ending call\n")
	doneChannel <- true
}

func main() {
	ready := make(chan bool)
	go longRunningCall(ready)
	done := <- ready
	fmt.Printf("call finished with %v\n", done)
}
	
