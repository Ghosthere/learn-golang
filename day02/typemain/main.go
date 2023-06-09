package main

import (
	"fmt"
	"time"

	"learn.org/types"
)

func testPoint() {
	fmt.Println(time.Now())
	var a, b = 5, 6
	c, d := types.PointChange(&a, b)
	fmt.Printf("c = %v, d = %v", c, d)
	fmt.Printf("\na = %v, b = %v", a, b)
}

func main() {
	a := types.StructTest(2, 3)
	fmt.Println(a.X)
	fmt.Printf("x: %v\n", a)

	p := &a
	p.X = 1e9
	fmt.Println(p.X)

	var (
		v1 = types.StructTest(1, 2) // has type Vertex
		v2 = types.StructTest(1, 0) // Y:0 is implicit
		v3 = types.StructTest1()    // X:0 and Y:0
		p1 = types.StructTest(1, 2) // has type *Vertex
		p2 = &p1
	)
	fmt.Println(v1, p2, v2, v3)
}
