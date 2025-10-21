package model

import (
	"time"
)

type Transaction struct {
	ID            string    `json:"id" gorm:"primaryKey"`
	UserID        string    `json:"user_id"`
	SeatID        string    `json:"seat_id"`
	ShowtimeID    string    `json:"showtime_id"`
	PaymentMethod string    `json:"payment_method"`
	TotalPrice    float64   `json:"total_price"`
	Status        string    `json:"status"`
	ExpiredAt     time.Time `json:"expired_at"`
	BookedAt      time.Time `json:"booked_at"`
}
