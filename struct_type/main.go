package main

import "fmt"

func main() {
	type Age int

	type Person struct {
		Name string
		Age  Age
	}

	fmt.Println(Person{Name: "Alice", Age: 30})
}
