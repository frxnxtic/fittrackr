package handlers

import (
	"Fittrackr/pkg/database"
	"Fittrackr/pkg/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func PostExerciseHandler(w http.ResponseWriter, r *http.Request) {
	var exercise models.Exercise
	err := json.NewDecoder(r.Body).Decode(&exercise)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	exerciseId := mux.Vars(r)["ID"]

	err = database.PostExercise(ctx, exerciseId, exercise)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Exercise created successfully"))

}

func GetExerciseHandler(w http.ResponseWriter, r *http.Request) {
	exerciseId := mux.Vars(r)["ID"]

	ctx := r.Context()
	exercise, err := database.GetExercise(ctx, exerciseId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(exercise)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
