package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Task struct {
	ID       int
	Title    string
	DueDate  string
	Category string
	IsDone   bool
	UserID   int
}

func (u User) print() {
	fmt.Println("User: ", u.ID, u.Email, u.Name)
}

var userStorage []User
var authenticatedUser *User // zero value for the pointer is nil
var taskStorage []Task

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

		// this is sanitization
		if authenticatedUser == nil {
			return
		}
	}

	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "register-user":
		registerUser()
	case "list-task":
		listTask()
	case "login":
		login()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("Command is not valid", command)
	}

}

func createTask() {

	scanner := bufio.NewScanner(os.Stdin)

	var title, duedate, category string

	fmt.Println("Please enter the task title")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("Please enter the task category")
	scanner.Scan()
	category = scanner.Text()

	fmt.Println("Please enter the task due date")
	scanner.Scan()
	duedate = scanner.Text()

	task := Task{
		ID:       len(taskStorage) + 1,
		Title:    title,
		DueDate:  duedate,
		Category: category,
		IsDone:   false,
		UserID:   authenticatedUser.ID,
	}

	taskStorage = append(taskStorage, task)

	fmt.Println("task: ", title, category, duedate)

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
		ID:       len(userStorage) + 1,
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
	}

}

func listTask() {
	for _, task := range taskStorage {
		if task.UserID == authenticatedUser.ID {
			fmt.Println(task)
		}
	}
}
