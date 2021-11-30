package main

import (
	"github.com/gorilla/mux"
	"iTechArt_Lab/task_3_Web/handlers"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	handlers.ReadJSON()

	r.HandleFunc("/students", handlers.GetStudents).Methods("GET")
	r.HandleFunc("/students/{id}", handlers.GetStudent).Methods("GET")
	r.HandleFunc("/grade/{class}", handlers.GetStudentsByClass).Methods("GET")
	r.HandleFunc("/students", handlers.CreateStudents).Methods("POST")
	r.HandleFunc("/students/{id}", handlers.UpdateStudents).Methods("PUT")
	r.HandleFunc("/students/{id}", handlers.DeleteStudents).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))

}
