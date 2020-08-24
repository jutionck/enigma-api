package userrepository

import "github.com/jutionck/enigma-bank-api/models"

type UserInterfaceRepository interface {
	GetAllUsers() ([]*models.User, error)
	GetByID(id int) (*models.User, error)
}
