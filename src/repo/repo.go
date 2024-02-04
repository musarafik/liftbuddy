package repo

import (
	"errors"
	"fmt"

	"github.com/musarafik/liftbuddy/src/structs"
)

type Repo interface {
	GetExercises(exercisesRepo structs.ExercisesRepo) (exercises []structs.Exercise)
	AddExercise(exercisesRepo structs.ExercisesRepo, newExercise structs.Exercise) (err error)
}

func GetExercises(exercisesRepo structs.ExercisesRepo) []structs.Exercise {
	return exercisesRepo.Exercises
}

func AddExercise(exercisesRepo *structs.ExercisesRepo, newExercise structs.Exercise) (err error) {
	err = nil
	exercises := exercisesRepo.Exercises

	if contains(exercises, newExercise) {
		err = errors.New("could not add new exercise since it already exists")
	} else {
		exercises = append(exercises, newExercise)
		fmt.Println(exercises)
		exercisesRepo.Exercises = exercises
	}

	return err
}

func contains(exercisesToSearch []structs.Exercise, exerciseToFind structs.Exercise) (doesContain bool) {
	doesContain = false

	for _, exercise := range exercisesToSearch {
		if exerciseToFind == exercise {
			doesContain = true
		}
	}

	return doesContain
}
