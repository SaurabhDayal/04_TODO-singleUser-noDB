package main

import (
	"04_TODO-SingleUser-noDB/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/todo", handlers.ReadAllTodo).Methods("GET")
	r.HandleFunc("/todo/{id}", handlers.ReadTodo).Methods("GET")
	r.HandleFunc("/todo", handlers.CreateTodo).Methods("POST")
	r.HandleFunc("/todo", handlers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todo/{id}", handlers.DeleteTodo).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
