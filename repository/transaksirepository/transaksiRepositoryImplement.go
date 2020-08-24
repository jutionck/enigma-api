package transaksirepository

import (
	"database/sql"
	"github.com/jutionck/enigma-bank-api/models"
	"github.com/jutionck/enigma-bank-api/utils"
)

type transactionRepo struct {
	db *sql.DB
}

func (t *transactionRepo) GetByUserOwnerID(id string) ([]*models.Transaksi, error) {
	var transactions []*models.Transaksi
	rows, err := t.db.Query(utils.GET_BY_ID_USER_OWNER, id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var each = new(models.Transaksi)
		if err := rows.Scan(&each.TransID, &each.UserOwnerID, &each.TransDate, &each.Destination, &each.Amount,
			&each.Description, &each.StatusDel); err != nil {
			return nil, err
		}
		transactions = append(transactions, each)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return transactions, nil
}

func (t *transactionRepo) GetAllTransactions() ([]*models.Transaksi, error) {
	var transactions []*models.Transaksi
	rows, err := t.db.Query(utils.GET_ALL_TRANSACTION)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var each = new(models.Transaksi)
		if err := rows.Scan(&each.TransID, &each.UserOwnerID, &each.TransDate, &each.Destination, &each.Amount,
			&each.Description, &each.StatusDel); err != nil {
			return nil, err
		}
		transactions = append(transactions, each)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return transactions, nil
}

func (t *transactionRepo) GetByID(id string) (*models.Transaksi, error) {
	var transaction = new(models.Transaksi)
	row := t.db.QueryRow(utils.GET_BY_ID_TRANSACTION, id)

	if err := row.Scan(&transaction.TransID, &transaction.UserOwnerID, &transaction.TransDate, &transaction.Destination,
		&transaction.Amount, &transaction.Description, &transaction.StatusDel); err != nil {
		return nil, err
	}
	return transaction, nil
}
func (t *transactionRepo) Store(trans *models.Transaksi) error {
	tx, err := t.db.Begin()

	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(utils.CREATE_TRANSACTION)

	if err != nil {
		return err
	}

	res, err := stmt.Exec(trans.UserOwnerID, trans.TransDate, trans.Destination,
		trans.Amount, trans.Description)

	if err != nil {
		return tx.Rollback()
	}

	_, err = res.LastInsertId()

	if err != nil {
		return tx.Rollback()
	}
	stmt.Close()
	return tx.Commit()
}

func NewTransactionRepo(db *sql.DB) TransaksiInterfaceRepo {
	return &transactionRepo{db}
}