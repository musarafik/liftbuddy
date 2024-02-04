package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/musarafik/liftbuddy/src/api/handlers"
)

func CreateServer() {
	fmt.Printf("Starting server...")
	http.HandleFunc("/get", handlers.GetExercisesHandler)
	http.HandleFunc("/create", handlers.CreateExerciseHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
