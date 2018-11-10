package main

import "fmt"

//
// type whatever struct { optional-fields }
// func (var *whatever) receivingFunction(arguments) returnvalues { ... }

type person struct {
	firstname string
	lastname string
}

func (p *person) print() {
	fmt.Printf("%s %s\n", p.firstname, p.lastname)
}

func main() {
	p := &person{"Harry", "Kodden"}
	p.print()
}
