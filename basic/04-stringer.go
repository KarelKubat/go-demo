package main

import "fmt"

type person struct {
	firstname string
	lastname string
}

// type Stringer interface {
//   String() string
// }
// --- used in %v expansion of Printf and family
// --- and in typecasts such as string(myVar)

func (p *person) String() string {
	return p.firstname + " " + p.lastname
}

func main() {
	fmt.Printf("%v\n", &person{
		firstname: "Harry",
		lastname: "Kodden",
	})
}
