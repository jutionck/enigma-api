package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jutionck/enigma-bank-api/delivery/middleware"
	"github.com/jutionck/enigma-bank-api/models"
	"github.com/jutionck/enigma-bank-api/usecase/transaksiusecase"
	"github.com/jutionck/enigma-bank-api/utils"
	"log"
	"net/http"
)

type TransaksiHandler struct {
	transUseCase transaksiusecase.TransaksiInterfaceUseCase
}

func TransRoute(r *mux.Router, service transaksiusecase.TransaksiInterfaceUseCase) {
	transHandler := TransaksiHandler{transUseCase: service}
	r.Use(middleware.ActivityLogMiddleware)
	s := r.PathPrefix("/transaksi").Subrouter()
	s.HandleFunc("", transHandler.ReturnAllFilms).Methods(http.MethodGet)
	s.HandleFunc("/{id}", transHandler.ReturnFindFilm).Methods(http.MethodGet)
	s.HandleFunc("/user/{id}", transHandler.ReturnFindTransaksiByUserID).Methods(http.MethodGet)
	s.HandleFunc("", transHandler.ReturnCreateFilm).Methods(http.MethodPost)
}

func (f *TransaksiHandler) ReturnAllFilms(w http.ResponseWriter, r *http.Request) {
	var response utils.Response
	rows, err := f.transUseCase.GetAllTransactions()
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

func (f *TransaksiHandler) ReturnFindFilm(w http.ResponseWriter, r *http.Request) {
	var response utils.Response
	w.Header().Set("Content-Type", "application/json")
	code := utils.DecodePathVariable("id", r)
	result, err := f.transUseCase.GetByID(code)

	response.Status = http.StatusOK
	response.Message = "Success: Film SelectByID"
	response.Result = result
	data, err := json.Marshal(response)

	if err != nil {
		log.Println(err)
		w.Write([]byte("Error occurred"))
	} else {
		w.Header().Set("content-type", "application/json")
		w.Write(data)
	}
}

func (f *TransaksiHandler) ReturnFindTransaksiByUserID(w http.ResponseWriter, r *http.Request) {
	var response utils.Response
	w.Header().Set("Content-Type", "application/json")
	code := utils.DecodePathVariable("id", r)
	result, err := f.transUseCase.GetByUserOwnerID(code)

	response.Status = http.StatusOK
	response.Message = "Success: Film SelectByID"
	response.Result = result
	data, err := json.Marshal(response)

	if err != nil {
		log.Println(err)
		w.Write([]byte("Error occurred"))
	} else {
		w.Header().Set("content-type", "application/json")
		w.Write(data)
	}
}

func (f *TransaksiHandler) ReturnCreateFilm(w http.ResponseWriter, r *http.Request) {
	var film models.Transaksi
	var response utils.Response
	err := json.NewDecoder(r.Body).Decode(&film) // json ke struct
	if err != nil {
		w.Write([]byte("Error"))
		log.Print(err)
	}

	err = f.transUseCase.Store(&film)
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