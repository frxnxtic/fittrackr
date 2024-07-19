package models

type WorkoutRecord struct {
	ID         string   `json:"id" firestore:"id"`                   // Идентификатор записи
	Date       string   `json:"date" firestore:"date"`               // Дата тренировки
	ExerciseID string   `json:"exercise_id" firestore:"exercise_id"` // Идентификатор упражнения
	Reps       int      `json:"reps" firestore:"reps"`               // Количество повторений в текущей тренировке
	Weight     float64  `json:"weight" firestore:"weight"`           // Рабочий вес в текущей тренировке
	Feelings   Feelings `json:"feelings" firestore:"feelings"`       // Ощущения по нагрузке
}
