package model

type Movie struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	Duration int    `json:"duration"`
	Synopsis string `json:"synopsis"`
}
