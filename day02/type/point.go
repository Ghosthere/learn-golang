package types

import "fmt"

func PointChange(a *int, b int) (int, int) {
	p := a          // point to a
	fmt.Println(*p) // read a through the pointer
	*p = 1
	fmt.Println(*a)
	b--
	return *p, b
}
