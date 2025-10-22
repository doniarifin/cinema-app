package model

type SeatTransaction struct {
	ID            string `json:"id" gorm:"primaryKey"`
	TransactionID string `json:"transaction_id"`
	SeatID        string `json:"seat_id"`
	ShowtimeID    string `json:"showtime_id"`
	Status        string `json:"status"`
	IsPaid        bool   `json:"is_paid"`
}
