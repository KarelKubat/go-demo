package main

import "fmt"

// Definition of behavior: any struct where makeNoise() can be called for,
// is a noiseMaker
type noiseMaker interface {
	// func keyword is implicit here
	makeNoise()
}

// Types can be specified as an interface as in argument n
func sayIt(n noiseMaker) {
	n.makeNoise()
}

// Empty struct and makeNoise() to make it a noiseMaker
type dog struct {}
func (d *dog) makeNoise() {
	fmt.Printf("woof\n")
}

// Another struct and makeNoise() to make it a noiseMaker
type cat struct{}
func (c *cat) makeNoise() {
	fmt.Printf("meow\n")
}

func main() {
	sayIt(&cat {})
	// Or:   c := &cat{}; sayIt(c)
	// Or:   c := cat{}; sayIt(&c)
	
	sayIt(&dog {})
}
