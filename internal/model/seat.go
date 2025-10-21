package model

type Seat struct {
	ID         string `json:"id" gorm:"primaryKey"`
	BranchID   string `json:"branch_id"`
	ShowtimeID string `json:"showtime_id"`
	SeatNumber string `json:"seat_number"`
	IsBooked   bool   `json:"is_booked" gorm:"default:false"`
}
