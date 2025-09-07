package service

import (
	"chin_server/internal/model"
	"chin_server/internal/repository"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Register(email, password, name string) (*model.User, error) {
	// 1. Kiểm tra email đã tồn tại chưa
	existing, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, fmt.Errorf("email already registered")
	}
	// 2. Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 3. Tạo user mới
	user := &model.User{
		Email:    email,
		Password: string(hashed),
		Username: name,
	}

	// 4. Lưu vào db
	err = s.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
