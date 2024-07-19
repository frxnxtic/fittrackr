package models

type Exercise struct {
	ID           string  `json:"id" firestore:"id"`                       // Идентификатор упражнения
	Name         string  `json:"name" firestore:"name"`                   // Название упражнения
	Description  string  `json:"description" firestore:"description"`     // Описание техники выполнения
	RestTime     int     `json:"rest_time" firestore:"rest_time"`         // Время на отдых в секундах
	Reps         int     `json:"reps" firestore:"reps"`                   // Рекомендуемое количество повторений
	TargetWeight float64 `json:"target_weight" firestore:"target_weight"` // Целевой рабочий вес
}
