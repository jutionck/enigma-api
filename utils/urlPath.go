package utils

import (
	"github.com/gorilla/mux"
	"net/http"
)

func DecodePathVariable(val string, r *http.Request) string {
	param := mux.Vars(r)
	return param[val]
}
