package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello to TODO app")

	command := flag.String("command", "no-command", "command to run")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	//input: name
	//input: category
	//input: due date
	if *command == "create-task" {
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

	} else if *command == "create-category" {
		var title, color string

		fmt.Println("Please enter the category title")
		scanner.Scan()
		title = scanner.Text()

		fmt.Println("Please enter the category color")
		scanner.Scan()
		color = scanner.Text()

		fmt.Println("category: ", title, color)
	} else if *command == "register-user" {
		var id, email, password string

		fmt.Println("Please enter user email")
		scanner.Scan()
		email = scanner.Text()

		fmt.Println("Please enter user password")
		scanner.Scan()
		password = scanner.Text()

		id = email

		fmt.Println("user: ", id, email, password)
	}

}
