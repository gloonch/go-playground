package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	var s Shape
	s = Circle{Radius: 3}
	fmt.Println(s.Area())

	// using type assertion
	describe(42)
	describe("name")
}

// func describe(value any) {
func describe(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Println("integer: ", v)
	case string:
		fmt.Println("string: ", v)
	case float64:
		fmt.Println("float64: ", v)
	default:
		fmt.Println("Unknown type")
	}
}
