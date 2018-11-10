package main

import (
	"fmt"
	"time"
)

func shoutAfter(d time.Duration) {
	time.Sleep(d)
	fmt.Printf("shouting after %v\n", d)
}

func main() {
	go shoutAfter(1 * time.Second)
	fmt.Printf("2 seconds until main() will exit\n");
	shoutAfter(2 * time.Second)
}
