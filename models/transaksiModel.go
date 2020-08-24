package models

type Transaksi struct {
	TransID     string `json:"trans_id"`
	UserOwnerID string    `json:"user_owner_id"`
	TransDate   string `json:"trans_date"`
	Destination string `json:"destination"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	StatusDel   int    `json:"-"`

}
