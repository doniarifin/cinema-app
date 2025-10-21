package repository

import (
	model "cinema-app/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindByEmail(string) (*model.User, error)
	FindByID(string) (*model.User, error)
	Create(*model.User) error
	Update(*model.User) error
	Delete([]string) error
}

type userRepositoryGorm struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryGorm{db}
}

func (u userRepositoryGorm) FindAll() ([]model.User, error) {
	var users []model.User
	result := u.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (u userRepositoryGorm) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := u.db.Where(&model.User{Email: email}).First(&user)
	if err.Error != nil {
		return nil, err.Error
	}
	return &user, nil
}

func (u userRepositoryGorm) FindByID(id string) (*model.User, error) {
	var user model.User
	err := u.db.Where(&model.User{ID: id}).First(&user)
	if err.Error != nil {
		return nil, err.Error
	}
	return &user, nil
}

func (u userRepositoryGorm) Create(m *model.User) error {
	return u.db.Create(m).Error
}

func (u userRepositoryGorm) Update(m *model.User) error {
	return u.db.Where(model.User{ID: m.ID}).Save(m).Error
}

func (u userRepositoryGorm) Delete(ids []string) error {
	users := model.User{}
	return u.db.Where("id IN ?", ids).Delete(&users).Error
}
