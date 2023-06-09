package data

import "fmt"

func ArrayCovert(a, b, c, d int) []int {
	x := []int{a, b, c, d}
	return x
}

func ArraySlice() {
	q := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	p := []bool{true, false, true, false, true, false}
	fmt.Printf("p: %v\n", p)

	s := []struct {
		x int
		y bool
	}{
		{1, true},
		{2, false},
		{3, true},
	}

	fmt.Println(s)
}
