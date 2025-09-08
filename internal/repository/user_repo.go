package repository

import (
	"chin_server/internal/model"

	"gorm.io/gorm"
)

// UserRepository: cung cấp các hàm thao tác với bảng users
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository: khởi tạo repository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser thêm user mới vào DB
func (r *UserRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

// GetUserByEmail, tìm user theo email
func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// get user bởi id
func (r *UserRepository) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
