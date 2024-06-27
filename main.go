package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strconv"
	"todocli/constant"
	"todocli/contract"
	"todocli/entity"
	"todocli/filestore"
)

type Task struct {
	ID         int
	Title      string
	DueDate    string
	CategoryID int
	IsDone     bool
	UserID     int
}

type Category struct {
	ID     int
	Title  string
	Color  string
	UserID int
}

var (
	userStorage       []entity.User
	authenticatedUser *entity.User // zero value for the pointer is nil
	taskStorage       []Task
	categoryStorage   []Category
	serializationMode string
)

const (
	userStoragePath = "user.txt"
)

func main() {

	sm := flag.String("serialize-mode", constant.SERIALIZATION_MODE_JSON, "serialization mode to write data to file")
	command := flag.String("command", "no-command", "command to run")
	flag.Parse()

	fmt.Println("Hello to TODO app")

	switch *sm {
	case constant.SERIALIZATION_MODE_MANDARAVARDI:
		serializationMode = constant.SERIALIZATION_MODE_MANDARAVARDI
	default:
		serializationMode = constant.SERIALIZATION_MODE_JSON
	}

	var userfileStore = filestore.New(userStoragePath, serializationMode)

	users := userfileStore.Load()
	userStorage = append(userStorage, users...)

	for {
		runCommand(userfileStore, *command)

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Please enter another command")
		scanner.Scan()
		*command = scanner.Text()
	}

}

func runCommand(store contract.UserWriteStore, command string) {
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
		registerUser(store)
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

	fmt.Println("Please enter the task category ID")
	scanner.Scan()
	category = scanner.Text()

	categoryID, err := strconv.Atoi(category)
	if err != nil {
		fmt.Printf("Category id is not valid integer, %v\n", err)

		return
	}

	isFound := false
	for _, c := range categoryStorage {
		if c.ID == categoryID && c.UserID == authenticatedUser.ID {
			isFound = true

			break
		}
	}

	if !isFound {
		fmt.Printf("Category id is not found\n")

		return
	}
	fmt.Println("Please enter the task due date")
	scanner.Scan()
	duedate = scanner.Text()

	task := Task{
		ID:         len(taskStorage) + 1,
		Title:      title,
		DueDate:    duedate,
		CategoryID: categoryID,
		IsDone:     false,
		UserID:     authenticatedUser.ID,
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

	category := Category{
		ID:     len(categoryStorage) + 1,
		Title:  title,
		Color:  color,
		UserID: authenticatedUser.ID,
	}

	categoryStorage = append(categoryStorage, category)

}

func registerUser(store contract.UserWriteStore) {
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

	user := entity.User{
		ID:       len(userStorage) + 1,
		Name:     name,
		Email:    email,
		Password: hashPassword(password),
	}
	userStorage = append(userStorage, user)

	//writeUserToFile(user)
	store.Save(user)
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
		if user.Email == email && user.Password == hashPassword(password) {
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

func hashPassword(password string) string {
	hash := md5.Sum([]byte(password))

	return hex.EncodeToString(hash[:])
}
