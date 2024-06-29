package filestore

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"todocli/constant"
	"todocli/entity"
)

type FileStore struct {
	filePath          string
	serializationMode string
}

// constructor
func New(path, serializationMode string) FileStore {
	return FileStore{filePath: path, serializationMode: serializationMode}
}

func (f FileStore) Save(u entity.User) {
	f.writeUserToFile(u)
}

func (f FileStore) Load() []entity.User {
	var uStorage []entity.User

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
		var userStruct = entity.User{}

		switch f.serializationMode {
		case constant.SERIALIZATION_MODE_MANDARAVARDI:
			var dErr error
			userStruct, dErr = deserializeFromMandaravardi(u)
			if dErr != nil {
				fmt.Println("Can't deserialize user record to user struct", dErr)

				return nil
			}
		case constant.SERIALIZATION_MODE_JSON:
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

		//uStorage = append(f, userStruct)
	}
	return uStorage
}

func (f FileStore) writeUserToFile(user entity.User) {

	// save user data in user.txt file
	// create user.txt file
	// write user record in the user.txt file

	var file *os.File

	_, err := os.Stat(f.filePath)
	if err != nil {
		fmt.Println("Path does not exist! ", err)

		var cErr error
		file, cErr = os.Create(f.filePath)
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
	if f.serializationMode == constant.SERIALIZATION_MODE_MANDARAVARDI {
		data = []byte(fmt.Sprintf("id: %d, name: %s, email: %s, password: %s\n",
			user.ID, user.Name, user.Email, user.Password))
	} else if f.serializationMode == constant.SERIALIZATION_MODE_JSON {
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

func deserializeFromMandaravardi(userStr string) (entity.User, error) {
	if userStr == "" {
		return entity.User{}, errors.New("User string is empty")
	}

	var user = entity.User{}

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

				return entity.User{}, errors.New("strconv error")
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
