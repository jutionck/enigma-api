package filmusecase

import (
	"github.com/jutionck/enigma-netflix-api/models"
	"github.com/jutionck/enigma-netflix-api/repository/filmrepository"
)

type filmUseCase struct {
	filmRepo filmrepository.FilmInterfaceRepo
}

func (f filmUseCase) GetAllFilms() ([]*models.Film, error) {
	films, err := f.filmRepo.GetAllFilms()
	if err != nil {
		return nil, err
	}
	return films, nil
}

func (f filmUseCase) PostFilm(film *models.Film) error {
	err := f.filmRepo.PostFilm(film)
	if err != nil {
		return err
	}
	return nil
}

func (f filmUseCase) GetFilmByID(filmID string) (models.Film, error) {
	film, err := f.filmRepo.GetFilmByID(filmID)
	if err != nil {
		return film, err
	}
	return film, nil
}

func NewFilmUseCase(filmRepo filmrepository.FilmInterfaceRepo) FilmInterfaceUseCase {
	return filmUseCase{filmRepo}
}