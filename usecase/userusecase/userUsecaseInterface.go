package userusecase

import "github.com/jutionck/enigma-bank-api/models"

type UserInterfaceUseCase interface {
	GetAllUsers() ([]*models.User, error)
	GetByID(id int) (*models.User, error)

}
