package model

type CinemaBranch struct {
	ID         string `json:"id" gorm:"primaryKey"`
	BranchName string `json:"branch_name"`
	City       string `json:"city"`
}
