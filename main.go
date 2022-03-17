package main

import (
	"fmt"

	"github.com/Saza-ku/go-genq/v1/pkg/genq"
)

type Person struct {
	FamilyName string
	FirstName  string
	Age        int
	Point      int
}

type Store struct {
	Id   int
	Name string
}

type Book struct {
	StoreId int
	Title   string
	Authors []string
}

func main() {
	s := &[]Person{
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
			return genq.Sum(genq.Select(func(p Person) int { return p.Point }, genq.From(&g.Group)))
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

	books := &[]Book{
		{0, "hoge", []string{"asano", "yoshida"}},
		{0, "hoge", []string{"kamiki"}},
		{2, "hoge", []string{"hamada", "matsumoto", "kamiki"}},
		{2, "hoge", []string{"kawai", "suzuki"}},
		{2, "hoge", []string{"harada", "asano", "himura", "matsumoto"}},
		{3, "hoge", []string{"ota", "furusawa"}},
		{5, "hoge", []string{"uchida", "kawai", "suzuki"}},
		{5, "hoge", []string{"matsumoto"}},
		{5, "hoge", []string{"asano", "hamada"}},
		{6, "hoge", []string{"kuwahara"}},
	}

	fmt.Println("===============================================================")

	fmt.Println(genq.SelectMany(
		func(b Book) []string { return b.Authors },
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
						func(b Book) []string { return b.Authors },
						genq.From(books),
					),
				),
			),
		).First()
	fmt.Println(a)

	fmt.Println("================================================================")

	stores := &[]Store{
		{0, "tokyo"},
		{1, "kyoto"},
		{2, "shiga"},
		{3, "osaka"},
		{4, "chiba"},
		{5, "okinawa"},
		{6, "hokkaido"},
	}

	b :=
		genq.Join(
			func(s Store, b Book) bool { return s.Id == b.StoreId },
			genq.From(stores),
			genq.From(books),
		).ToSlice()
	fmt.Println(b)

	fmt.Println("================================================================")
	cs := []int{1, 2, 3, 4, 5}
	cq := genq.From(&cs)
	cs = append(cs, 6)
	fmt.Println(genq.Aggregate(func(s, x int) int { return s + x }, cq))
}
