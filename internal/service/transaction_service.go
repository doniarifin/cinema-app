package service

import (
	"cinema-app/internal/model"
	"cinema-app/internal/repository"
	"cinema-app/internal/utils"
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

func (s *TransactionService) CreateTransaction(userID, showtimeID string, seats []*model.SeatTransaction, payment string) (*model.Transaction, error) {
	// total := 0.0

	var showtime model.Showtime
	if err := s.db.First(&showtime, "id = ?", showtimeID).Error; err != nil {
		return nil, fmt.Errorf("showtime not found: %v", err)
	}
	total := float64(len(seats)) * float64(showtime.Price)

	t := model.Transaction{
		ID:            utils.GenerateUUID(),
		UserID:        userID,
		ShowtimeID:    showtimeID,
		Showtime:      showtime,
		Seats:         seats,
		PaymentMethod: payment,
		TotalPrice:    total,
		Status:        "pending",
		ExpiredAt:     time.Now().Add(10 * time.Minute),
		BookedAt:      time.Now(),
	}

	if err := s.repo.Create(&t); err != nil {
		return nil, err
	}

	for _, seat := range seats {
		seat.TransactionID = t.ID
		seat.ShowtimeID = t.ShowtimeID
		if err := s.repo.AddSeat(seat); err != nil {
			return nil, err
		}
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

	if t.Status == "paid" || t.Status == "canceled" {
		return fmt.Errorf("failed, current status transaction: %s", t.Status)
	}

	return s.repo.UpdateStatus(id, "paid")
}

func (s *TransactionService) CancelOrder(id string) error {
	t, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if t.Status == "canceled" {
		return fmt.Errorf("failed, current status transaction: %s", t.Status)
	}

	return s.repo.UpdateStatus(id, "canceled")
}
