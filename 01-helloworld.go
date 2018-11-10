package main

import "fmt"

// Lowercase 'a' in 'add': won't get exported out of this package
// Combined argument type: func add(a int, b int) int { ... }
func add(a, b int) int {
	return a + b
}

// You can return more than 1 thing
func addAndSubtract(a, b int) (int, int) {
	return a+b, a-b
}

func main() {
	fmt.Printf("Hello World\n");
	fmt.Printf("3+7 is %v\n", add(3, 7))

	sum, diff := addAndSubtract(12, 3)
	fmt.Printf("12+3 is %d, 12-3 is %d\n", sum, diff)
}
