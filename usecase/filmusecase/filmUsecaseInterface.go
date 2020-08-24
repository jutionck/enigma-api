package filmusecase

import "github.com/jutionck/enigma-netflix-api/models"

type FilmInterfaceUseCase interface {
	GetAllFilms() ([]*models.Film, error)
	GetFilmByID(string) (models.Film, error)
	PostFilm(film *models.Film) error
}
