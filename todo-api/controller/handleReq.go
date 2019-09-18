package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Punam-Gaikwad/Getgoing/todo-api/model"
	"github.com/Punam-Gaikwad/Getgoing/todo-api/views"
)

var (
	er      error
	records []views.PostRequest
)

func handleReq() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			if err := model.CreateTodo(data.Name, data.Todo); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodGet {
			name := r.URL.Query().Get("name")
			if name != "" {
				records, er = model.ReadByName(name)
			} else {
				records, er = model.ReadAll()
			}
			if er != nil {
				w.Write([]byte(er.Error()))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(records)
		} else if r.Method == http.MethodDelete {
			name := r.URL.Path[1:]
			if err := model.DeleteTodo(name); err != nil {
				w.Write([]byte("Some error"))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(struct {
				Status string `json:"status"`
			}{"Item deleted"})
		}
	}
}
