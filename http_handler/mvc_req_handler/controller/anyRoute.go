package controller

import "net/http"

func AnyRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("In AnyRoute handler /"))
}
