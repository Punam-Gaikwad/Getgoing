package model

import (
	"fmt"

	"github.com/Punam-Gaikwad/Getgoing/todo-api/views"
)

//CreateTodo -
func CreateTodo(name, todo string) error {
	insertQ, err := con.Query("INSERT INTO TODO VALUES(?,?)", name, todo)
	defer insertQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//ReadAll -
func ReadAll() ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM TODO")
	if err != nil {
		return nil, err
	}

	todos := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}

	return todos, nil
}

//ReadByName -
func ReadByName(name string) ([]views.PostRequest, error) {
	rows, err := con.Query("SELECT * FROM TODO WHERE name=?", name)
	if err != nil {
		return nil, err
	}

	todos := []views.PostRequest{}
	for rows.Next() {
		data := views.PostRequest{}
		rows.Scan(&data.Name, &data.Todo)
		todos = append(todos, data)
	}

	return todos, nil
}

//DeleteTodo -
func DeleteTodo(name string) error {
	deleteQ, err := con.Query("DELETE FROM TODO WHERE name=?", name)
	defer deleteQ.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
