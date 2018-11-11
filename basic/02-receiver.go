package main

import "fmt"

// Methods: functions specific for a struct. The struct: receiver.
// type whatever struct { optional-fields }
// func (var *whatever) receivingFunction(arguments) returnvalues { ... }

// ----------------------------------------------- typical person example
type person struct {
	firstname string
	lastname string
}

func (p *person) print() {
	fmt.Printf("%s %s\n", p.firstname, p.lastname)
}

func testPerson() {
	p := &person{"Harry", "Kodden"}
	p.print()
}

// -------------------------------------------------- non-struct example
// Receivers don't have to be structs, though they most often are.
// Methods that modify their receiver, must have a pointer-to receiver type.

type prime int

func (p *prime) isPrime() bool {
	for i := 2; i < int(*p); i++ {
		if int(*p) % i == 0 {
			return false
		}
	}
	return true
}

func (p *prime) next() {
	for {
		(*p)++
		if p.isPrime() {
			return
		}
	}
}

func testPrime() {
	p := prime(1)
	for int(p) < 100 {
		p.next()
		fmt.Printf("%v ", int(p))
	}
	fmt.Printf("\n")
}

// ---------------------------------------------------------------- Test all
func main() {
	testPerson()
	testPrime()
}
