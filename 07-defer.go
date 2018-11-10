package main

import "fmt"

func testFunc() {
	defer fmt.Printf("done\n")	
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", i)
	}
}

func main() {
	testFunc()
}
