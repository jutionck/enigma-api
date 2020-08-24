package utils

const (
	SELECT_ALL_FILM = `SELECT id_film,film_title,film_duration,film_image_url,film_synopsis,film_status FROM film`
	SELECT_FILM_BYID = `SELECT id_film,film_title,film_duration,film_image_url,film_synopsis,film_status FROM film WHERE id_film =?`
	SQL_INSERT_FILM = `INSERT INTO film(film_title,film_duration,film_image_url,film_synopsis) VALUES (?,?,?,?)`
)
