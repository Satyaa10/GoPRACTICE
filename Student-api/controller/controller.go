package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	// Import the models package
	"github.com/Satyaa10/GOPRACTICE/Student-api/models"
	"github.com/Satyaa10/GOPRACTICE/Student-api/storage"
	"github.com/go-chi/chi/v5"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	// this empty  variable which  will store the json data which is commig
	var student models.Student

	//&student  in this save the variable in student , it uses pointer because it needs to update the student varicble
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	student.ID = storage.NextID
	storage.Students = append(storage.Students, student)
	storage.NextID++
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)

}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(storage.Students)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)
	var updated models.Student
	json.NewDecoder(r.Body).Decode(&updated)

	for i, s := range storage.Students {
		if s.ID == id {
			s.Name = updated.Name
			s.Age = updated.Age
			s.Email = updated.Email
			storage.Students[i] = s
			json.NewEncoder(w).Encode(s)
			return
		}

	}
	http.Error(w, "student not found ", http.StatusNotFound)

}

func GetStudentByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)
	for _, s := range storage.Students {
		if s.ID == id {
			json.NewEncoder(w).Encode(s)
			return
		}
	}
	http.Error(w, "student not found", http.StatusNotFound)

}
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)
	for i, s := range storage.Students {
		if s.ID == id {
			storage.Students = append(storage.Students[:i], storage.Students[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Students  not found ", http.StatusNotFound)
}
