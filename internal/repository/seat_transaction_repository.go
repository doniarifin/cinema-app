package repository

import (
	"cinema-app/internal/dto"
	"cinema-app/internal/model"

	"gorm.io/gorm"
)

type SeatTransactionRepo interface {
	GetAll() ([]model.SeatTransaction, error)
	GetByID(id string) (*model.SeatTransaction, error)
	GetsSeatTrx(f dto.SeatTrx) (*[]model.SeatTransaction, error)
	Create(c *model.SeatTransaction) error
	Update(c *model.SeatTransaction) error
	UpdateMany(trxId string, c *model.SeatTransaction) error
	Delete(id string) error
}

type seatTrxRepo struct {
	db *gorm.DB
}

func NewSeatTransactionRepo(db *gorm.DB) SeatTransactionRepo {
	return &seatTrxRepo{db}
}

func (r *seatTrxRepo) GetAll() ([]model.SeatTransaction, error) {
	var movies []model.SeatTransaction
	err := r.db.Find(&movies).Error
	return movies, err
}

func (r *seatTrxRepo) GetByID(id string) (*model.SeatTransaction, error) {
	var c model.SeatTransaction
	if err := r.db.First(&c, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *seatTrxRepo) GetsSeatTrx(filter dto.SeatTrx) (*[]model.SeatTransaction, error) {
	var seats []model.SeatTransaction
	query := r.db.Model(&model.SeatTransaction{})

	if filter.ID != "" {
		query = query.Where("id = ?", filter.ID)
	}
	if filter.IDs != nil {
		query = query.Where("id IN ?", filter.IDs)
	}
	if filter.TransactionID != "" {
		query = query.Where("transaction_id = ?", filter.TransactionID)
	}
	if filter.ShowtimeID != "" {
		query = query.Where("showtime_id = ?", filter.ShowtimeID)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	if filter.IsPaid != nil {
		query = query.Where("is_paid = ?", *filter.IsPaid)
	}

	err := query.Find(&seats).Error
	return &seats, err

}

func (r *seatTrxRepo) Create(c *model.SeatTransaction) error {
	return r.db.Create(c).Error
}

func (r *seatTrxRepo) Update(c *model.SeatTransaction) error {
	return r.db.Save(c).Error
}

func (r *seatTrxRepo) UpdateMany(trxId string, updated *model.SeatTransaction) error {
	updateData := map[string]any{
		"status":  updated.Status,
		"is_paid": updated.IsPaid,
	}

	return r.db.Model(&model.SeatTransaction{}).
		Where("transaction_id = ?", trxId).
		Updates(updateData).Error
}

func (r *seatTrxRepo) Delete(id string) error {
	return r.db.Delete(&model.SeatTransaction{}, id).Error
}
