package handler

import (
	"Fittrackr/pkg/database"
	"Fittrackr/pkg/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ExerciseHandler struct{}

func (h ExerciseHandler) PostHandler(w http.ResponseWriter, r *http.Request) {
	var exercise model.Exercise
	err := json.NewDecoder(r.Body).Decode(&exercise)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	err = database.PostModel(ctx, exercise, "exercises")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Exercise created successfully"))

}

func (h ExerciseHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	exerciseId := mux.Vars(r)["ID"]

	ctx := r.Context()
	exercise, err := database.GetModel[model.Exercise](ctx, exerciseId, "exercises")
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

func (h ExerciseHandler) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	exercises, err := database.GetAllModels[model.Exercise](ctx, "exercises")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(exercises)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (h ExerciseHandler) PutHandler(w http.ResponseWriter, r *http.Request) {
	var exercise model.Exercise
	err := json.NewDecoder(r.Body).Decode(&exercise)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	exerciseId := mux.Vars(r)["ID"]

	err = database.UpdateModel(ctx, exerciseId, exercise, "exercises")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Exercise updated successfully"))
}

func (h ExerciseHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	exerciseId := mux.Vars(r)["ID"]

	ctx := r.Context()
	err := database.DeleteModel(ctx, exerciseId, "exercises")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Exercise deleted successfully"))
}
