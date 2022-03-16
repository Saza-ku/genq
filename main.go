package main

import (
	"fmt"
	"genq/pkg/genq"
)

func main() {
	s := []int{1, 2, 3, 4}
	q := genq.From(s)
	fmt.Println(q.ToSlice())
}
