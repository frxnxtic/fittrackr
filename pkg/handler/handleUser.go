package handler

import (
	"Fittrackr/pkg/database"
	"Fittrackr/pkg/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type UserHandler struct{}

func (h UserHandler) PostHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	err = database.PostModel(ctx, user, "users")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully!"))
}

func (h UserHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["ID"]

	ctx := r.Context()
	user, err := database.GetModel[model.User](ctx, userId, "users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (h UserHandler) GetAllHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := database.GetAllModels[model.User](ctx, "users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (h UserHandler) PutHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	userId := mux.Vars(r)["ID"]

	err = database.UpdateModel(ctx, userId, user, "users")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("user updated successfully"))
}

func (h UserHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["ID"]

	ctx := r.Context()
	err := database.DeleteModel(ctx, userId, "users")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}
