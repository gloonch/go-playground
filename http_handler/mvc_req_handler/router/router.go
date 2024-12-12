package router

import (
	"fmt"
	"go-playground/mvc_req_handler/controller"
	"net/http"
)

func Init() {
	http.HandleFunc("/", controller.AnyRoute)
	http.HandleFunc("/getData", controller.GetDataRoute)
	http.HandleFunc("/getData/{id}", controller.GetDataWithIdRoute)

	fmt.Println("Server is running on port 8000")
	http.ListenAndServe(":8000", nil)
}
