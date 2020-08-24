package models

type Film struct {
	FilmID     		string   	`json:"film_id"`
	FilmTitle     	string   	`json:"film_title"`
	FilmDuration    string   	`json:"film_duration"`
	FilmImageURL 	string 		`json:"film_image_url"`
	FilmSynopsis    string      `json:"film_synopsis"`
	FilmStatus      string        `json:"film_status"`

}
