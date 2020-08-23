package utils

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Result    interface{} `json:"result"`
}
