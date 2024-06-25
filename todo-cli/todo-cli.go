package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

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

func (u User) print() {
	fmt.Println("User: ", u.ID, u.Email, u.Name)
}

var (
	userStorage       []User
	authenticatedUser *User // zero value for the pointer is nil
	taskStorage       []Task
	categoryStorage   []Category
	serializationMode string
)

const (
	userStoragePath                 = "user.txt"
	SERIALIZATION_MODE_JSON         = "json"
	SERIALIZATION_MODE_MANDARAVARDI = "mandaravardi"
)

var userfileStore = fileStore{
	filePath: userStoragePath,
}

func main() {

	sm := flag.String("serialize-mode", SERIALIZATION_MODE_JSON, "serialization mode to write data to file")
	command := flag.String("command", "no-command", "command to run")
	flag.Parse()

	loadUserFromStorage(userfileStore, *sm)

	fmt.Println("Hello to TODO app")

	switch *sm {
	case SERIALIZATION_MODE_MANDARAVARDI:
		serializationMode = SERIALIZATION_MODE_MANDARAVARDI
	default:
		serializationMode = SERIALIZATION_MODE_JSON
	}
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
		registerUser(userfileStore)
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

type userWriteStore interface {
	Save(u User)
}

type userReadStore interface {
	Load(serializeMode string) []User
}

func registerUser(store userWriteStore) {
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

func loadUserFromStorage(store userReadStore, serializeMode string) {
	users := store.Load(serializeMode)

	userStorage = append(userStorage, users...)
}

func loadUserStorageFromFile(serializeMode string) {

}

func deserializeFromMandaravardi(userStr string) (User, error) {
	if userStr == "" {
		return User{}, errors.New("User string is empty")
	}

	var user = User{}

	userFields := strings.Split(userStr, ",")
	for _, field := range userFields {
		values := strings.Split(field, ": ")
		if len(values) != 2 {
			fmt.Println("Record is not valid, skipping... ", len(values))

			continue
		}
		fieldName := strings.ReplaceAll(values[0], " ", "")
		fieldValue := values[1]

		switch fieldName {
		case "id":
			id, err := strconv.Atoi(fieldValue)
			if err != nil {
				fmt.Println("strconv error ", err)

				return User{}, errors.New("strconv error")
			}
			user.ID = id
		case "name":
			user.Name = fieldValue
		case "email":
			user.Email = fieldValue
		case "password":
			user.Password = fieldValue
		}
	}

	return user, nil
}

func (f fileStore) writeUserToFile(user User) {

	// save user data in user.txt file
	// create user.txt file
	// write user record in the user.txt file

	var file *os.File

	_, err := os.Stat(userStoragePath)
	if err != nil {
		fmt.Println("Path does not exist! ", err)

		var cErr error
		file, cErr = os.Create(userStoragePath)
		if err != nil {
			fmt.Println("Can't create the user.txt file", cErr)

			return
		}

	} else {
		var oErr error
		file, oErr = os.OpenFile(f.filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
		if err != nil {
			fmt.Println("Can't create or open file ", oErr)

			return
		}
		defer file.Close()
	}

	var data []byte
	// serializing the user struct
	if serializationMode == SERIALIZATION_MODE_MANDARAVARDI {
		data = []byte(fmt.Sprintf("id: %d, name: %s, email: %s, password: %s\n",
			user.ID, user.Name, user.Email, user.Password))
	} else if serializationMode == SERIALIZATION_MODE_JSON {
		data, err = json.Marshal(user)
		if err != nil {
			fmt.Println("Can't marshal user struct to json ", err)

			return
		}

		data = append(data, []byte("\n")...)
	} else {
		fmt.Println("Invalid serialization mode")

		return
	}

	numberOfWrittenBytes, wErr := file.Write(data)
	if wErr != nil {
		fmt.Printf("Can't write to the file %v ", wErr)

		return
	}

	fmt.Println("Number of written bytes ", numberOfWrittenBytes)

	// call defer functions

}

func hashPassword(password string) string {
	hash := md5.Sum([]byte(password))

	return hex.EncodeToString(hash[:])
}

type fileStore struct {
	filePath string
}

func (f fileStore) Save(u User) {
	f.writeUserToFile(u)
}

func (f fileStore) Load(serializationMode string) []User {
	var uStorage []User

	file, err := os.Open(f.filePath)
	if err != nil {
		fmt.Println("Can't open the file ", err)
	}

	var data = make([]byte, 10240)
	_, oErr := file.Read(data)
	if oErr != nil {
		fmt.Println("Can't read from the file ", oErr)

		return nil
	}

	var dataStr = string(data)

	dataStr = strings.Trim(dataStr, "\n")

	userSlice := strings.Split(dataStr, "\n")

	for _, u := range userSlice {
		var userStruct = User{}

		switch serializationMode {
		case SERIALIZATION_MODE_MANDARAVARDI:
			var dErr error
			userStruct, dErr = deserializeFromMandaravardi(u)
			if dErr != nil {
				fmt.Println("Can't deserialize user record to user struct", dErr)

				return nil
			}
		case SERIALIZATION_MODE_JSON:
			if u[0] != '{' && u[len(u)-1] != '}' {
				continue
			}
			uErr := json.Unmarshal([]byte(u), &userStruct)
			if uErr != nil {
				fmt.Println("Can't deserialize user record to user struct with json mode ", uErr)

				return nil
			}
		default:
			fmt.Println("invalid serialization mode")

			return nil
		}

		uStorage = append(userStorage, userStruct)
	}
	return uStorage
}
