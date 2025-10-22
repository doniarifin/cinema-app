package dto

type Seat struct {
	ID         string   `form:"id"`
	IDs        []string `form:"ids"`
	BranchID   string   `form:"branch_id"`
	ShowtimeID string   `form:"showtime_id"`
	SeatNumber string   `form:"seat_number"`
	Status     string   `form:"status"`
	IsBooked   *bool    `form:"is_booked"`
}
