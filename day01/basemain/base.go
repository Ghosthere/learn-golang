package main

import (
	"fmt"
	"math"
	"runtime"

	b "learn.org/base"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("windows")
	}
}

func main1() {
	// fmt.Println("Beijing = ", BEIJING)
	fmt.Println(b.City())
	s := b.SHANGHAI
	fmt.Println(s)
	fmt.Printf("math.Sqrt(2): %v\n", b.Sqrt(2))
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
