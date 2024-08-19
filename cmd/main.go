package main

import (
	"Fittrackr/pkg/database"
	"Fittrackr/pkg/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Initialize Firestore database
	err := database.InitFirestore()
	if err != nil {
		log.Fatalf("Failed to initialize Firestore: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// Create an instance of ExerciseHandler
	exerciseHandler := &handler.ExerciseHandler{}
	userHandler := &handler.UserHandler{}
	workoutHandler := &handler.WorkoutHandler{}

	// Create a new router
	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/exercise", exerciseHandler.PostHandler).Methods("POST")
	apiRouter.HandleFunc("/exercises", exerciseHandler.GetAllHandler).Methods("GET")
	apiRouter.HandleFunc("/exercise/{ID:[a-zA-Z0-9_-]+}", exerciseHandler.GetHandler).Methods("GET")
	apiRouter.HandleFunc("/exercise/{ID:[a-zA-Z0-9_-]+}", exerciseHandler.PutHandler).Methods("PUT")
	apiRouter.HandleFunc("/exercise/{ID:[a-zA-Z0-9_-]+}", exerciseHandler.DeleteHandler).Methods("DELETE")

	apiRouter.HandleFunc("/user", userHandler.PostHandler).Methods("POST")
	apiRouter.HandleFunc("/users", userHandler.GetAllHandler).Methods("GET")
	apiRouter.HandleFunc("/user/{ID:[a-zA-Z0-9_-]+}", userHandler.GetHandler).Methods("GET")
	apiRouter.HandleFunc("/user/{ID:[a-zA-Z0-9_-]+}", userHandler.PutHandler).Methods("PUT")
	apiRouter.HandleFunc("/user/{ID:[a-zA-Z0-9_-]+}", userHandler.DeleteHandler).Methods("DELETE")

	apiRouter.HandleFunc("/workout", workoutHandler.PostHandler).Methods("POST")
	apiRouter.HandleFunc("/workouts", workoutHandler.GetAllHandler).Methods("GET")
	apiRouter.HandleFunc("/workout/{ID:[a-zA-Z0-9_-]+}", workoutHandler.GetHandler).Methods("GET")
	apiRouter.HandleFunc("/workout/{ID:[a-zA-Z0-9_-]+}", workoutHandler.PutHandler).Methods("PUT")
	apiRouter.HandleFunc("/workout/{ID:[a-zA-Z0-9_-]+}", workoutHandler.DeleteHandler).Methods("DELETE")

	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
