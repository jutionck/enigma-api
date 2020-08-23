package filmrepository

import (
	"database/sql"
	"github.com/jutionck/enigma-netflix-api/models"
	"github.com/jutionck/enigma-netflix-api/utils"
	"log"
)

type FilmRepository struct {
	Conn *sql.DB
}

func (f *FilmRepository) GetAllFilms() ([]*models.Film, error) {
	var films []*models.Film
	rows, err := f.Conn.Query(utils.SELECT_ALL_FILM)
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		film := models.Film{}
		if err := rows.Scan(&film.FilmID, &film.FilmTitle, &film.FilmDuration,&film.FilmImageURL,&film.FilmSynopsis, &film.FilmStatus); err != nil {
			log.Fatalf("%v", err)
			return nil, err
		} else {
			films = append(films, &film)
		}
	}
	return films, nil
}

func (f *FilmRepository) GetFilmByID(filmID string) (models.Film, error) {
	var film models.Film
	query := utils.SELECT_FILM_BYID
	err := f.Conn.QueryRow(query, filmID).Scan(&film.FilmID, &film.FilmTitle, &film.FilmDuration,&film.FilmImageURL,&film.FilmSynopsis,&film.FilmStatus)
	if err != nil {
		return film, err
	}
	return film, nil
}

func (f *FilmRepository) PostFilm(film *models.Film) error {
	tx, err := f.Conn.Begin()
	query := utils.SQL_INSERT_FILM
	stmt, err := tx.Exec(query, film.FilmTitle, film.FilmDuration, film.FilmImageURL, film.FilmSynopsis)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	_, err = stmt.LastInsertId()


	if err != nil {
		log.Fatal(err)
	}
	return tx.Commit()
}

func NewFilmRepository(db *sql.DB) FilmInterfaceRepo {
	return &FilmRepository{Conn: db}
}
