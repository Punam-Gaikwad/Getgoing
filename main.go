package main

import (
	"fmt"
	"net/http"
)

func main() {
	//Request Handler
	mux := http.NewServeMux()

	//Handles request coming to endpoint - localhost:3030/
	//works as Router
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request recieved")
		fmt.Println(r.Method)
		w.Write([]byte("Hello World\n"))
	})

	//Handles request coming to endpoint - localhost:3030/getgoing
	mux.HandleFunc("/getgoing", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Udemy GoLang getgoing course\n"))
	})

	//connect to the API listening on port 3030
	http.ListenAndServe("localhost:3030", mux)

	//steps to execute are
	//1. go run main.go
	//2. curl -i localhost:3030/ or
	// 	 curl -i localhost:3030/getgoing

	//can also send POST req as - curl -X POST -i localhost:3030/
}
