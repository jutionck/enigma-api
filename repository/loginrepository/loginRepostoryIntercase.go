package loginrepository

import "github.com/jutionck/enigma-bank-api/models"

type LoginInterfaceRepository interface {
	GetAllUsers() ([]*models.Login, error)
	GetByID(id int) (*models.Login, error)
}
