package userrepository

import (
	"database/sql"
	"github.com/jutionck/enigma-bank-api/models"
	"github.com/jutionck/enigma-bank-api/utils"
)

type userRepo struct {
	db *sql.DB
}

func (u *userRepo) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	rows, err := u.db.Query(utils.GET_ALL_USER)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var each = new(models.User)
		if err := rows.Scan(&each.UserID, &each.LoginOwnerID, &each.Balance, &each.StatusDel); err != nil {
			return nil, err
		}
		users = append(users, each)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepo) GetByID(id int) (*models.User, error) {
	var user = new(models.User)
	row := u.db.QueryRow(utils.GET_BY_ID_LOGIN_OWNER, id)

	if err := row.Scan(&user.UserID, &user.LoginOwnerID, &user.Balance, &user.StatusDel); err != nil {
		return nil, err
	}
	return user, nil
}

func NewUserRepo(db *sql.DB) UserInterfaceRepository {
	return &userRepo{db}
}
