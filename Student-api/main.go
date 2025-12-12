package main

import (
	"fmt"
	"net/http"

	"github.com/Satyaa10/GOPRACTICE/Student-api/router"
)

func main() {
	r := router.NewRouter()
	fmt.Println("student api to manage ")

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)

}
