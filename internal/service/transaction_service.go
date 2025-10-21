package service

import (
	"cinema-app/internal/model"
	"cinema-app/internal/repository"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type TransactionService struct {
	repo *repository.TransactionRepository
	db   *gorm.DB
}

func NewTransactionService(repo *repository.TransactionRepository, db *gorm.DB) *TransactionService {
	return &TransactionService{
		repo: repo,
		db:   db,
	}
}

func (s *TransactionService) CreateTransaction(userID, showtimeID string, seats string, payment string) (*model.Transaction, error) {
	// total := 0.0

	var showtime model.Showtime
	if err := s.db.First(&showtime, "id = ?", showtimeID).Error; err != nil {
		return nil, fmt.Errorf("showtime not found: %v", err)
	}
	total := float64(showtime.Price)

	t := model.Transaction{
		UserID:        userID,
		ShowtimeID:    showtimeID,
		SeatID:        seats,
		PaymentMethod: payment,
		TotalPrice:    total,
		Status:        "pending",
		ExpiredAt:     time.Now().Add(10 * time.Minute),
	}

	if err := s.repo.Create(&t); err != nil {
		return nil, err
	}

	SeatTrx := []model.SeatTransaction{}

	for _, seat := range SeatTrx {
		seat.TransactionID = t.ID
		_ = s.repo.AddSeat(&seat)
	}

	return &t, nil
}

func (s *TransactionService) ExpirePendingTransactions() error {
	txs, err := s.repo.FindExpiredTransactions()
	if err != nil {
		return err
	}
	for _, t := range txs {
		_ = s.repo.ExpireTransaction(t.ID)
	}
	return nil
}

func (s *TransactionService) MarkAsPaid(id string) error {
	t, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if t.Status != "pending" {
		return fmt.Errorf("transaction is not pending, current status: %s", t.Status)
	}

	return s.repo.UpdateStatus(id, "paid")
}
