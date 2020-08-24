package transaksiusecase

import (
	"github.com/jutionck/enigma-bank-api/models"
	"github.com/jutionck/enigma-bank-api/repository/transaksirepository"
	"github.com/jutionck/enigma-bank-api/utils"
)

type transactionUseCase struct {
	transactionRepo transaksirepository.TransaksiInterfaceRepo
}

func (t *transactionUseCase) GetByUserOwnerID(id string) ([]*models.Transaksi, error) {
	res, err := t.transactionRepo.GetByUserOwnerID(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *transactionUseCase) GetAllTransactions() ([]*models.Transaksi, error) {
	res, err := t.transactionRepo.GetAllTransactions()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *transactionUseCase) GetByID(id string) (*models.Transaksi, error) {
	res, err := t.transactionRepo.GetByID(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *transactionUseCase) Store(trans *models.Transaksi) error {
	err := utils.ValidateInputNotEmpty(trans.UserOwnerID, trans.TransDate, trans.Destination,
		trans.Amount, trans.Description)

	if err != nil {
		return err
	}

	if err = t.transactionRepo.Store(trans); err != nil {
		return err
	}

	return nil
}

func NewTransactionUseCase(repo transaksirepository.TransaksiInterfaceRepo) TransaksiInterfaceUseCase {
	return &transactionUseCase{repo}
}