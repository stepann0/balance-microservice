package models

import (
	"time"
)

// Главный счет (аккаунт) пользователя
type Account struct {
	ID      uint    `json:"id" gorm:"primary_key"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

// Резервация денег пользователя перед транзакцией
type ReservationAccount struct {
	ID        uint    `json:"id"`
	ServiceID uint    `json:"service_id"`
	Amount    float64 `json:"balance"`
}

// Все поддерживаемые виды услуг
type Service struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

// Все совершенные транзакции
type Transaction struct {
	ID     uint      `json:"id" gorm:"primary_key"`
	Amount float64   `json:"amount"`
	Time   time.Time `json:"time"`
	Status bool      `json:"status"`
	Error  string    `json:"error"`
}

type Payment struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	ServiceID uint      `json:"service_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"time"`
}
