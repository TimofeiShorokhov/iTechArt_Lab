package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Student struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Class string `json:"class"`
}

var Students []Student

func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Students)
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, i := range Students {
		if i.ID == params["id"] {
			json.NewEncoder(w).Encode(i)
			return
		}
	}
	json.NewEncoder(w).Encode(&Student{})
}

func GetStudentsByClass(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, i := range Students {
		if i.Class == params["class"] {
			json.NewEncoder(w).Encode(i)
			return
		}
	}
	json.NewEncoder(w).Encode(&Student{})
}

func CreateStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student Student
	_ = json.NewDecoder(r.Body).Decode(&student)
	newId := len(Students) + 1
	for _, i := range Students {
		if i.ID == strconv.Itoa(newId) {
			newId += 1
		}
	}
	//student.ID = strconv.Itoa(rand.Intn(1000000))
	student.ID = strconv.Itoa(newId)
	Students = append(Students, student)
	json.NewEncoder(w).Encode(student)
}

func UpdateStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, i := range Students {
		if i.ID == params["id"] {
			Students = append(Students[:index], Students[index+1:]...)
			var student Student
			_ = json.NewDecoder(r.Body).Decode(&student)
			student.ID = params["id"]
			Students = append(Students, student)
			json.NewEncoder(w).Encode(student)
			return
		}
	}
	json.NewEncoder(w).Encode(Students)
}
func DeleteStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, i := range Students {
		if i.ID == params["id"] {
			Students = append(Students[:index], Students[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Students)
}

func ReadJSON() []Student {
	filename, err := os.Open("task_3_Web/mocks/students.json")
	if err != nil {
		log.Fatal(err)
	}

	defer filename.Close()
	data, err := ioutil.ReadAll(filename)

	if err != nil {
		log.Fatal(err)
	}

	jsonErr := json.Unmarshal(data, &Students)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return Students
}
