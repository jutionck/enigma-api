package delivery

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/jutionck/enigma-bank-api/repository/transaksirepository"
	"github.com/jutionck/enigma-bank-api/usecase/transaksiusecase"
)

func Init(r *mux.Router, db *sql.DB) {
	filmRepo := transaksirepository.NewTransactionRepo(db)
	filmUseCase := transaksiusecase.NewTransactionUseCase(filmRepo)
	TransRoute(r, filmUseCase)
}