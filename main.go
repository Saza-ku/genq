package main

import (
	"fmt"
	"genq/pkg/genq"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	s := []*Person{{"yamada", 13}, {"suzuki", 29}, {"yoshida", 41}, {"tanaja", 9}}

	x := genq.Select(
		func(p *Person) string { return p.Name },
		genq.Where(
			func(p *Person) bool { return p.Age >= 20 },
			genq.From(s)),
	).ToSlice()
	fmt.Println(x)
}
