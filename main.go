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
	q := genq.From(s)
	q1 := genq.Where(q, func(p *Person) bool { return p.Age >= 20 })
	q2 := genq.Select(q1, func(p *Person) string { return p.Name })
	fmt.Println(q2.ToSlice())
}
