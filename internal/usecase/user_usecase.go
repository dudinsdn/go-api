package usecase

import (
	"my/go-api/internal/model"
	"my/go-api/internal/repository"
)

type UserUsecase interface {
	FetchUsers() ([]model.User, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) FetchUsers() ([]model.User, error) {
	return u.repo.GetAll()
}
