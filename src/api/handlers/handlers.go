package handlers

import (
	"fmt"
	"net/http"

	"github.com/musarafik/liftbuddy/src/repo"
	"github.com/musarafik/liftbuddy/src/structs"
)

func GetExercisesHandler(w http.ResponseWriter, r *http.Request) {
	exercises := repo.GetExercises()

	fmt.Fprint(w, exercises)
}

func CreateExerciseHandler(w http.ResponseWriter, r *http.Request) {
	newExercise := structs.Exercise{
		Name:    "pull up",
		NumReps: 5,
		NumSets: 5,
		Weight:  10,
	}

	err := repo.AddExercise(newExercise)
	if err != nil {
		fmt.Print(err)
		fmt.Fprint(w, "error")
	} else {
		fmt.Fprint(w, newExercise)
	}
}
