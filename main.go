package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/musarafik/structs"
)

func createExerciseHandler(w http.ResponseWriter, r *http.Request) {
	newExercise := structs.Exercise{
		Name:    "pull up",
		NumReps: 5,
		NumSets: 5,
		Weight:  10,
	}

	fmt.Fprint(w, newExercise)
}

func main() {
	fmt.Printf("Hello world")

	http.HandleFunc("/create", createExerciseHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
