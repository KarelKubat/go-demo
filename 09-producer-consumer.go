package main

import "fmt"

func producer(c chan int64) {
	var a int64 = 1
	var b int64 = 2
	for {
		c <- a
		a, b = b, a+b
	}
}

func consumer(c chan int64, done chan bool) {
	for {
		val := <- c
		fmt.Printf("%v ", val)
		if val > 100000000 {
			fmt.Printf("\n")
			done <- true
			return
		}
	}
}

func main() {
	numberChannel := make(chan int64, 3) // note: size=3, buffered channel
	doneChannel := make(chan bool)

	go producer(numberChannel)
	go consumer(numberChannel, doneChannel)

	<- doneChannel
}
