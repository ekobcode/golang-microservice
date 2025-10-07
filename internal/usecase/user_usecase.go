package usecase

import (
	"golang-microservice/internal/entity"
	"golang-microservice/internal/repository"
)

type UserUsecase interface {
	Create(user *entity.User) error
	GetAll() ([]entity.User, error)
	GetByID(id int64) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id int64) error
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo}
}

func (u *userUsecase) Create(user *entity.User) error {
	return u.repo.Create(user)
}

func (u *userUsecase) GetAll() ([]entity.User, error) {
	return u.repo.FindAll()
}

func (u *userUsecase) GetByID(id int64) (*entity.User, error) {
	return u.repo.FindByID(id)
}

func (u *userUsecase) Update(user *entity.User) error {
	return u.repo.Update(user)
}

func (u *userUsecase) Delete(id int64) error {
	return u.repo.Delete(id)
}
