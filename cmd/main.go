package main

import (
	"Fittrackr/pkg/database"
	"Fittrackr/pkg/handler"
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

	exerciseHandler := &handler.ExerciseHandler{}

	r := mux.NewRouter()
	r.HandleFunc("/exercise", exerciseHandler.PostHandler).Methods("POST")
	r.HandleFunc("/exercises", exerciseHandler.GetAllHandler).Methods("GET")
	r.HandleFunc("/exercise/{ID:[a-zA-Z0-9_-]+}", exerciseHandler.GetHandler).Methods("GET")
	r.HandleFunc("/exercise/{ID:[a-zA-Z0-9_-]+}", exerciseHandler.PutHandler).Methods("PUT")
	r.HandleFunc("/exercise/{ID:[a-zA-Z0-9_-]+}", exerciseHandler.DeleteHandler).Methods("DELETE")

	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
