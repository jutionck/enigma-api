package loginusecase

import "github.com/jutionck/enigma-bank-api/models"

type LoginInterfaceUsecase interface {
	GetAllUsers() ([]*models.Login, error)
	GetByID(id int) (*models.Login, error)
}
