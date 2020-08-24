package models

type Login struct {
	LoginID   int    `json:"login_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	StatusDel int    `json:"-"`
}

