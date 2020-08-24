package userusecase

import (
	"github.com/jutionck/enigma-bank-api/models"
	"github.com/jutionck/enigma-bank-api/repository/userrepository"
)

type userUseCase struct {
	userRepo userrepository.UserInterfaceRepository
}

func (u *userUseCase) GetAllUsers() ([]*models.User, error) {
	res, err := u.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *userUseCase) GetByID(id int) (*models.User, error) {
	res, err := u.userRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewUserUseCase(repo userrepository.UserInterfaceRepository) UserInterfaceUseCase {
	return &userUseCase{repo}
}