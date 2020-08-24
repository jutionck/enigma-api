package transaksirepository

import "github.com/jutionck/enigma-bank-api/models"

type TransaksiInterfaceRepo interface {
	GetAllTransactions() ([]*models.Transaksi, error)
	GetByID(string) (*models.Transaksi, error)
	GetByUserOwnerID(string) ([]*models.Transaksi, error)
	Store(trans *models.Transaksi) error
}
