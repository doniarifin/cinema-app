package model

import "time"

type Showtime struct {
	ID       string    `json:"id" gorm:"primaryKey"`
	BranchID string    `json:"branch_id"`
	MovieID  string    `json:"movie_id"`
	DateTime time.Time `json:"date_time"`
	Price    int       `json:"price"`
}
