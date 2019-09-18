package controller

import (
	"net/http"
)

//Register - Check API health
//Request  - GET:/ping
//Response - pong
func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping())
	mux.HandleFunc("/", handleReq())
	return mux
}
