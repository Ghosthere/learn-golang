package types

type Vertex struct {
	X int
	Y int
}

func StructTest(x, y int) Vertex {
	return Vertex{x, y}
}

func StructTest1() Vertex {
	return Vertex{}
}
