package tests

import (
	"testing"

	"github.com/musarafik/liftbuddy/src/repo"
	"github.com/musarafik/liftbuddy/src/structs"
)

// Adding just one new exercise to an empty repo
func TestAddSingleExercise(t *testing.T) {
	var exercisesRepo structs.ExercisesRepo
	var exercises []structs.Exercise

	newExercise := structs.Exercise{
		Name:    "bench press",
		NumReps: 5,
		NumSets: 5,
		Weight:  10,
	}

	err := repo.AddExercise(&exercisesRepo, newExercise)
	if err != nil {
		t.Error("already exists")
	}

	exercises = repo.GetExercises(exercisesRepo)

	if exercises[0] != newExercise {
		t.Error("error")
	}
}

// Adding a couple new exercises to an empty repo
func TestAddMultipleExercises(t *testing.T) {
	var exercisesRepo structs.ExercisesRepo
	var newExercises []structs.Exercise

	newExercise1 := structs.Exercise{
		Name:    "bench press",
		NumReps: 5,
		NumSets: 5,
		Weight:  10,
	}
	newExercise2 := structs.Exercise{
		Name:    "push ups",
		NumReps: 5,
		NumSets: 5,
		Weight:  10,
	}
	newExercise3 := structs.Exercise{
		Name:    "sit ups",
		NumReps: 5,
		NumSets: 5,
		Weight:  10,
	}
	newExercise4 := structs.Exercise{
		Name:    "squat",
		NumReps: 5,
		NumSets: 5,
		Weight:  10,
	}
	newExercise5 := structs.Exercise{
		Name:    "deadlift",
		NumReps: 5,
		NumSets: 5,
		Weight:  10,
	}

	newExercises = append(newExercises, newExercise1)
	newExercises = append(newExercises, newExercise2)
	newExercises = append(newExercises, newExercise3)
	newExercises = append(newExercises, newExercise4)
	newExercises = append(newExercises, newExercise5)

	for _, newExercise := range newExercises {
		err := repo.AddExercise(&exercisesRepo, newExercise)
		if err != nil {
			t.Error("exercise already exists")
		}
	}

	for i, exercise := range repo.GetExercises(exercisesRepo) {
		if exercise != newExercises[i] {
			t.Error("error")
		}
	}
}

// Adding duplicate exercises to repo
func TestAddDuplicateExercises(t *testing.T) {
	var exercisesRepo structs.ExercisesRepo
	var newExercises []structs.Exercise

	newExercise1 := structs.Exercise{
		Name:    "bench press",
		NumReps: 5,
		NumSets: 5,
		Weight:  10,
	}
	newExercise2 := structs.Exercise{
		Name:    "bench press",
		NumReps: 5,
		NumSets: 5,
		Weight:  10,
	}

	newExercises = append(newExercises, newExercise1)
	newExercises = append(newExercises, newExercise2)

	for _, newExercise := range newExercises {
		err := repo.AddExercise(&exercisesRepo, newExercise)
		if err != nil {
			t.Logf("exercise already exists")
		}
	}

	if len(repo.GetExercises(exercisesRepo)) != 1 {
		t.Error("error")
	}
}
