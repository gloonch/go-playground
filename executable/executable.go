package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	executable, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(executable)
	fmt.Println("Executable binary path: ", exPath)

	time.Sleep(time.Second * 10)
}
