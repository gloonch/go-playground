package main

import "fmt"

func dangerous(task func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic happened: %v", r)
		}
	}()
	task()
	return nil
}

func main() {
	err := dangerous(func() {
		panic("panic to test")
	})
	if err != nil {
		fmt.Println("Recovered:", err)
	} else {
		fmt.Println("همه چیز خوب بود.")
	}
}
