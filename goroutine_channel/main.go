package main

import (
	"fmt"
	"time"
)

func sendStringData(ch chan string) {
	ch <- "hello from goroutine"
}

func main() {

	ch := make(chan string)
	go sendStringData(ch)

	msg := <-ch
	fmt.Println("given data: ", msg)
	time.Sleep(1 * time.Second)

}
