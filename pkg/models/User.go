package models

type User struct {
	ID      string   `json:"id" firestore:"id"`           // Идентификатор пользователя
	Name    string   `json:"name" firestore:"name"`       // Имя пользователя
	Height  float64  `json:"height" firestore:"height"`   // Рост в сантиметрах
	Weight  float64  `json:"weight" firestore:"weight"`   // Вес в килограммах
	Records []Record `json:"records" firestore:"records"` // Личные рекорды
}
