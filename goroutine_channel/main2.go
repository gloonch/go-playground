package main

import (
	"fmt"
)

func sendData(channel chan int) {
	for i := 0; i < 5; i++ {
		channel <- i
	}

}

func main() {

	mChannel := make(chan int)
	go sendData(mChannel)

	for i := 0; i < 5; i++ {
		fmt.Println(<-mChannel)
	}
}
