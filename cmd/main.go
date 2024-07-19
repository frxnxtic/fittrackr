package main

import (
	"Fittrackr/pkg/database"
	"Fittrackr/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	err := database.InitFirestore()
	if err != nil {
		log.Fatalf("Failed to initialize Firestore: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	r := mux.NewRouter()
	r.HandleFunc("/exercise/{ID:[a-zA-Z0-9_-]+}", handlers.PostExerciseHandler).Methods("POST")
	r.HandleFunc("/exercise/{ID:[a-zA-Z0-9_-]+}", handlers.GetExerciseHandler).Methods("GET")

	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
