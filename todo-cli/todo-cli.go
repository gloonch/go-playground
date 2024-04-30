package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func (u User) print() {
	fmt.Println("User: ", u.ID, u.Email, u.Name)
}

var userStorage []User
var authenticatedUser *User // zero value for the pointer is nil

func main() {
	fmt.Println("Hello to TODO app")

	command := flag.String("command", "no-command", "command to run")
	flag.Parse()

	for {
		runCommand(*command)

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Please enter another command")
		scanner.Scan()
		*command = scanner.Text()
	}

}

func runCommand(command string) {
	if command != "register-user" && command != "exit" && authenticatedUser == nil {
		login()
	}

	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()
	case "login":
		login()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("Command is not valid", command)
	}

}

func createTask() {

	if authenticatedUser != nil {
		authenticatedUser.print()
	}

	scanner := bufio.NewScanner(os.Stdin)

	var name, duedate, category string

	fmt.Println("Please enter the task title")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("Please enter the task category")
	scanner.Scan()
	category = scanner.Text()

	fmt.Println("Please enter the task due date")
	scanner.Scan()
	duedate = scanner.Text()

	fmt.Println("task: ", name, category, duedate)

}

func createCategory() {
	scanner := bufio.NewScanner(os.Stdin)

	var title, color string

	fmt.Println("Please enter the category title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("Please enter the category color")
	scanner.Scan()
	color = scanner.Text()

	fmt.Println("category: ", title, color)

}

func registerUser() {
	scanner := bufio.NewScanner(os.Stdin)

	var id, name, email, password string

	fmt.Println("Please enter user email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("Please enter user name")
	scanner.Scan()
	name = scanner.Text()

	fmt.Println("Please enter user password")
	scanner.Scan()
	password = scanner.Text()

	id = email

	fmt.Println("user: ", id, email, password)

	user := User{
		ID:       strconv.Itoa(len(userStorage) + 1),
		Name:     name,
		Email:    email,
		Password: password,
	}
	userStorage = append(userStorage, user)

}

func login() {
	scanner := bufio.NewScanner(os.Stdin)

	var id, email, password string

	fmt.Println("Please enter user email")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("Please enter user password")
	scanner.Scan()
	password = scanner.Text()

	fmt.Println("user: ", id, email, password)

	for _, user := range userStorage {
		if user.Email == email && user.Password == password {
			authenticatedUser = &user

			break
		}
	}

	// If there is a user record with corresponding data, then allow the user to continue
	if authenticatedUser == nil {
		fmt.Println("The email or password is not correct")
		return
	}

}
