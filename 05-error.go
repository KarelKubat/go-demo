package main

import "fmt"

// type error interface {
//   Error() string
// }

type divideByZeroError struct {
	numerator int
}

func newDivideByZeroError(a, b int) *divideByZeroError {
	return &divideByZeroError{
		numerator: a,
	}
}

func (d *divideByZeroError) Error() string {
	return fmt.Sprintf("cannot divide %d by 0", d.numerator)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, newDivideByZeroError(a, b)
	}
	return a/b, nil
}

func main() {
	result, err := divide(12, 2)
	fmt.Printf("result: %v, error: %v\n", result, err)

	// result, err are already known, so = instead of :=
	result, err = divide(10, 0)
	fmt.Printf("result: %v, error: %v\n", result, err)
}
