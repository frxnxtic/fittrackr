package handler

import (
	"Fittrackr/pkg/database"
	"Fittrackr/pkg/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ExerciseHandler struct{}

// PostHandler handles the HTTP POST request for creating a new exercise.
// It decodes the request body into an Exercise struct, validates it, and saves it to the database.
// If the request body is invalid or there is an error saving to the database, it returns an appropriate HTTP error response.
// If the exercise is created successfully, it returns a HTTP status code 201 (Created) with a success message.
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

// GetHandler handles the HTTP GET request for retrieving an exercise by ID.
// It takes the HTTP response writer and request as parameters.
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

// GetAllHandler handles the HTTP request to retrieve all exercises.
// It fetches all exercise models from the database and returns them as a JSON response.
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

// PutHandler updates an existing exercise with the provided exercise data.
//
// It decodes the request body into an Exercise struct and updates the exercise in the database.
// The exercise ID is extracted from the request URL using mux.Vars.
// If the decoding or updating process encounters an error, an appropriate HTTP error response is sent.
// If the update is successful, it returns a 200 OK status code and a success message.
func (h ExerciseHandler) PutHandler(w http.ResponseWriter, r *http.Request) {
	// PutHandler updates an existing exercise with the provided exercise data.

	// Decode the request body into an Exercise struct
	var exercise model.Exercise
	err := json.NewDecoder(r.Body).Decode(&exercise)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	exerciseId := mux.Vars(r)["ID"]

	// Update the exercise in the database
	err = database.UpdateModel(ctx, exerciseId, exercise, "exercises")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Exercise updated successfully"))
}

// DeleteHandler handles the HTTP DELETE request to delete an exercise.
// It takes the exercise ID from the request URL and deletes the corresponding exercise from the database.
// If the deletion is successful, it returns a 200 OK response with a success message.
// If there is an error during the deletion process, it returns a 500 Internal Server Error response.
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
