package handler

import (
	"Fittrackr/pkg/database"
	"Fittrackr/pkg/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type WorkoutHandler struct{}

// PostHandler handles the HTTP POST request for creating a new Workout.
// It decodes the request body into an Workout struct, validates it, and saves it to the database.
// If the request body is invalid or there is an error saving to the database, it returns an appropriate HTTP error response.
// If the Workout is created successfully, it returns a HTTP status code 201 (Created) with a success message.
func (h WorkoutHandler) PostHandler(w http.ResponseWriter, r *http.Request) {
	var workout model.WorkoutRecord
	err := json.NewDecoder(r.Body).Decode(&workout)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	err = database.PostModel(ctx, workout, "workouts")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Workout created successfully"))

}

// GetHandler handles the HTTP GET request for retrieving an Workout by ID.
// It takes the HTTP response writer and request as parameters.
func (h WorkoutHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	workoutId := mux.Vars(r)["ID"]

	ctx := r.Context()
	workout, err := database.GetModel[model.WorkoutRecord](ctx, workoutId, "workouts")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(workout)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// GetAllHandler handles the HTTP request to retrieve all Workouts.
// It fetches all Workout models from the database and returns them as a JSON response.
func (h WorkoutHandler) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	workouts, err := database.GetAllModels[model.WorkoutRecord](ctx, "workouts")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(workouts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// PutHandler updates an existing Workout with the provided Workout data.
//
// It decodes the request body into an Workout struct and updates the Workout in the database.
// The Workout ID is extracted from the request URL using mux.Vars.
// If the update is successful, it returns a 200 OK status code and a success message.
// If the decoding or updating process encounters an error, an appropriate HTTP error response is sent.
func (h WorkoutHandler) PutHandler(w http.ResponseWriter, r *http.Request) {
	// PutHandler updates an existing Workout with the provided Workout data.

	// Decode the request body into an Workout struct
	var workout model.WorkoutRecord
	err := json.NewDecoder(r.Body).Decode(&workout)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	WorkoutId := mux.Vars(r)["ID"]

	// Update the Workout in the database
	err = database.UpdateModel(ctx, WorkoutId, workout, "workouts")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Workout updated successfully"))
}

// DeleteHandler handles the HTTP DELETE request to delete an Workout.
// It takes the Workout ID from the request URL and deletes the corresponding Workout from the database.
// If the deletion is successful, it returns a 200 OK response with a success message.
// If there is an error during the deletion process, it returns a 500 Internal Server Error response.
func (h WorkoutHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	workoutId := mux.Vars(r)["ID"]

	ctx := r.Context()
	err := database.DeleteModel(ctx, workoutId, "workouts")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Workout deleted successfully"))
}
