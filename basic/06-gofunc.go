package main

import (
	"fmt"
	"time"
)

func testFunc(d time.Duration) {
	time.Sleep(d)
	fmt.Printf("testFunc ends after %v\n", d)
}

func main() {
	// go namedFunction(args)
	go testFunc(1 * time.Second)

	// go closure(args) { .... } (parameters)
	go func(d time.Duration) {
		time.Sleep(d)
		fmt.Printf("closure ends after %v\n", d);
	}(time.Second)
	
	fmt.Printf("main will stop in 2 seconds\n");
	time.Sleep(2 * time.Second)
}
