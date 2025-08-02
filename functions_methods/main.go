package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {

	// sum function
	sum := add(1, 2)
	fmt.Println("Sum result: ", sum)

	// division function with multiple return values
	division, err := divide(22, 4)
	if err != nil {
		fmt.Println("Error while dividing.")

		return
	}
	fmt.Println("Division result: ", division)

	// method (receiver function)
	user := User{Name: "Clay"}
	fmt.Println("Method call result: ", user.Greet())
	user.setName("Abdullah")
	fmt.Println("Method call result after changing the name: ", user.Greet())
}

type User struct {
	Name string
}

func (u User) Greet() string {
	return "Hello, " + u.Name
}

func (u *User) setName(name string) {
	u.Name = name
}
