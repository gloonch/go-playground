package controller

import "net/http"

func GetDataRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("In GetDataRoute handler /"))
}
