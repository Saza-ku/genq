package main

import (
	"fmt"
	"genq/pkg/genq"
)

type Person struct {
	FamilyName string
	FirstName  string
	Age        int
	Point      int
}

type Book struct {
	id      int
	title   string
	authors []string
}

func main() {
	s := []Person{
		{"kazuki", "yamada", 13, 27}, {"hanako", "suzuki", 29, 43}, {"kengo", "yoshida", 41, 82}, {"momoka", "tanaka", 9, 23},
		{"ryoya", "yamada", 83, 90}, {"takumi", "suzuki", 92, 43}, {"fumi", "yoshida", 8, 62}, {"ayumi", "tanaka", 12, 89},
		{"kei", "yamada", 63, 10}, {"hibiki", "suzuki", 31, 43}, {"fuka", "yoshida", 47, 97}, {"nanami", "tanaka", 120, 71},
	}

	x :=
		genq.Select(
			func(p Person) string { return p.FirstName + " " + p.FamilyName },
			genq.Where(
				func(p Person) bool { return 20 <= p.Age && p.Age <= 65 },
				genq.From(s)),
		).ToSlice()
	fmt.Println(x)

	y :=
		genq.Select(func(g genq.Group[string, Person]) int {
			return genq.Sum(genq.Select(func(p Person) int { return p.Point }, genq.From(g.Group)))
		},
			genq.GroupBy(
				func(p Person) string { return p.FamilyName },
				genq.From(s),
			),
		).ToSlice()
	fmt.Println(y)

	z :=
		genq.OrderBy(func(p Person) int { return p.Age },
			genq.From(s),
		).ToSlice()
	fmt.Println(z)

	books := []Book{
		{0, "hoge", []string{"asano", "yoshida"}},
		{0, "hoge", []string{"kamiki"}},
		{0, "hoge", []string{"hamada", "matsumoto", "kamiki"}},
		{0, "hoge", []string{"kawai", "suzuki"}},
		{0, "hoge", []string{"harada", "asano", "himura", "matsumoto"}},
		{0, "hoge", []string{"ota", "furusawa"}},
		{0, "hoge", []string{"uchida", "kawai", "suzuki"}},
		{0, "hoge", []string{"matsumoto"}},
		{0, "hoge", []string{"asano", "hamada"}},
		{0, "hoge", []string{"kuwahara"}},
	}

	fmt.Println("===============================================================")

	fmt.Println(genq.SelectMany(
		func(b Book) []string { return b.authors },
		genq.From(books),
	).ToSlice())

	a, _ :=
		genq.Select(
			func(g genq.Group[string, string]) string { return g.Key },
			genq.OrderByDescending(
				func(g genq.Group[string, string]) int { return len(g.Group) },
				genq.GroupBy(
					func(a string) string { return a },
					genq.SelectMany(
						func(b Book) []string { return b.authors },
						genq.From(books),
					),
				),
			),
		).First()
	fmt.Println(a)
}
