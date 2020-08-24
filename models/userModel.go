package models

type User struct {
	UserID       int `json:"user_id"`
	LoginOwnerID int `json:"login_owner_id"`
	Balance      int `json:"balance"`
	StatusDel    int `json:"status_del"`
}
