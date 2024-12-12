package controller

import "net/http"

func GetDataWithIdRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("In GetDataRoute handler /"))
}
