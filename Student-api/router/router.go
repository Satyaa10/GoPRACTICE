package router

import (
	"github.com/Satyaa10/GOPRACTICE/Student-api/controller"
	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/students", func(r chi.Router) {
		r.Get("/", controller.GetAllStudents)
		r.Post("/", controller.CreateStudent)
		r.Get("/{id}", controller.GetStudentByID)
		r.Put("/{id}", controller.UpdateStudent)
		r.Delete("/{id}", controller.DeleteStudent)
	})

	return r
}
