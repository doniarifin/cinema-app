package repository

import (
	"cinema-app/internal/model"
	"cinema-app/internal/utils"
	"time"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db}
}

func (r *TransactionRepository) Create(t *model.Transaction) error {
	return r.db.Create(t).Error
}

func (r *TransactionRepository) FindAll() ([]model.Transaction, error) {
	var txs []model.Transaction
	err := r.db.Find(&txs).Error
	return txs, err
}

func (r *TransactionRepository) FindByID(id string) (*model.Transaction, error) {
	var t model.Transaction
	if err := r.db.First(&t, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *TransactionRepository) UpdateStatus(id string, status string) error {
	return r.db.Model(&model.Transaction{}).Where("id = ?", id).Update("status", status).Error
}

func (r *TransactionRepository) AddSeat(seat *model.SeatTransaction) error {
	if seat.ID == "" {
		seat.ID = utils.GenerateUUID()
	}
	return r.db.Create(seat).Error
}

func (r *TransactionRepository) FindExpiredTransactions() ([]model.Transaction, error) {
	var txs []model.Transaction
	err := r.db.Where("status = ? AND expired_at <= ?", "pending", time.Now()).Find(&txs).Error
	return txs, err
}

func (r *TransactionRepository) ExpireTransaction(id string) error {
	return r.db.Model(&model.Transaction{}).Where("id = ?", id).Update("status", "expired").Error
}
