package main

import (
	"fmt"
	"math"

	b "learn.org/base"
)

func main() {
	// fmt.Println("Beijing = ", BEIJING)
	fmt.Println(b.City())
	s := b.SHANGHAI
	fmt.Println(s)
	fmt.Printf("math.Sqrt(2): %v\n", math.Sqrt(2))
}

func main2() {
	sum := 0
	for i := 1; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	for sum < 1000 {
		sum += sum
		if v := math.Pow(3, 5); v < float64(sum) {
			fmt.Printf("v: %v\n", v)
		}
	}

	fmt.Printf("sum: %v\n", sum)
}
