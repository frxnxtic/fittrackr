package models

type Record struct {
	ExerciseID string  `json:"exercise_id" firestore:"exercise_id"` // Идентификатор упражнения
	MaxWeight  float64 `json:"max_weight" firestore:"max_weight"`   // Максимальный рабочий вес
	MaxReps    int     `json:"max_reps" firestore:"max_reps"`       // Максимальное количество повторений
	Date       string  `json:"date" firestore:"date"`               // Дата рекорда
}
