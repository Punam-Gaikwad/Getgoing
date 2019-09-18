package main

import (
	"fmt"
	"net/http"

	"github.com/Punam-Gaikwad/Getgoing/todo-api/controller"
	"github.com/Punam-Gaikwad/Getgoing/todo-api/model"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving")
	http.ListenAndServe(":3000", mux)
}
