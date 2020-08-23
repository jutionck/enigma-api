package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jutionck/enigma-netflix-api/delivery/middleware"
	"github.com/jutionck/enigma-netflix-api/models"
	"github.com/jutionck/enigma-netflix-api/usecase/filmusecase"
	"github.com/jutionck/enigma-netflix-api/utils"
	"log"
	"net/http"
)

type FilmHandler struct {
	filmUseCase filmusecase.FilmInterfaceUseCase
}

func FilmRoute(r *mux.Router, service filmusecase.FilmInterfaceUseCase) {
	filmHandler := FilmHandler{filmUseCase: service}
	r.Use(middleware.ActivityLogMiddleware)
	s := r.PathPrefix("/film").Subrouter()
	s.HandleFunc("", filmHandler.ReturnAllFilms).Methods(http.MethodGet)
	s.HandleFunc("/{film_id}", filmHandler.ReturnFindFilm).Methods(http.MethodGet)
	s.HandleFunc("", filmHandler.ReturnCreateFilm).Methods(http.MethodPost)
}

func (f *FilmHandler) ReturnAllFilms(w http.ResponseWriter, r *http.Request) {
	var response utils.Response
	rows, err := f.filmUseCase.GetAllFilms()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}

	response.Status = http.StatusOK
	response.Message = "Success: Film Select"
	response.Result = rows

	//change data (rows) to JSON
	byteOfFilm, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Oops, Something when wrong !!"))
	}
	// set header look
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfFilm)
}

func (f *FilmHandler) ReturnFindFilm(w http.ResponseWriter, r *http.Request) {
	//var responseStud utils.Response
	w.Header().Set("Content-Type", "application/json")
	code := utils.DecodePathVariable("film_id", r)
	result, err := f.filmUseCase.GetFilmByID(code)
	data, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error occurred"))
	} else {
		w.Header().Set("content-type", "application/json")
		w.Write(data)
	}
}

func (f *FilmHandler) ReturnCreateFilm(w http.ResponseWriter, r *http.Request) {
	var film models.Film
	var response utils.Response
	err := json.NewDecoder(r.Body).Decode(&film) // json ke struct
	if err != nil {
		w.Write([]byte("Error"))
		log.Print(err)
	}

	err = f.filmUseCase.PostFilm(&film)
	if err != nil {
		w.Write([]byte(err.Error()))
		log.Print(err)
		return
	}
	response.Status = http.StatusOK
	response.Message = "Success: Film Insert !"
	response.Result = film

	byteOfFilm, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfFilm)
}