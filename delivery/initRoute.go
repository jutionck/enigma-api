package delivery

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/jutionck/enigma-netflix-api/repository/filmrepository"
	"github.com/jutionck/enigma-netflix-api/usecase/filmusecase"
)

func Init(r *mux.Router, db *sql.DB) {
	filmRepo := filmrepository.NewFilmRepository(db)
	filmUseCase := filmusecase.FilmInterfaceUseCase(filmRepo)
	FilmRoute(r, filmUseCase)
}