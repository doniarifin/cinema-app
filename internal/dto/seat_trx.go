package dto

type SeatTrx struct {
	ID            string   `form:"id"`
	IDs           []string `form:"ids"`
	TransactionID string   `form:"transaction_id"`
	SeatID        string   `form:"seat_id"`
	ShowtimeID    string   `form:"showtime_id"`
	Status        string   `form:"status"`
	IsPaid        *bool    `form:"is_paid"`
}
