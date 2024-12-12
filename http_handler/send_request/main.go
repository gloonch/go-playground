package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	httpClient := http.Client{Timeout: 1 * time.Nanosecond}

	request, _ := http.NewRequest(http.MethodGet, "www.google.com", nil)

	_, err := httpClient.Do(request)
	if os.IsTimeout(err) {
		fmt.Println("Timed out: ", err)
	}
}
