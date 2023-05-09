package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nataliadiasa/register/domain"
	"github.com/nataliadiasa/register/service"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var dat domain.Person
	if err := json.Unmarshal(body, &dat); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	h.service.Create(dat)
	w.WriteHeader(http.StatusCreated)
}

func (h Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	persons := h.service.GetAll()
	body, err := json.Marshal(persons)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = w.Write(body)
	if err != nil {
		panic(err)
	}
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	person, err := h.service.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Person doesn't exist"))
		return
	}

	body, err := json.Marshal(person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = w.Write(body)
	if err != nil {
		panic(err)
	}
}

func (h Handler) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi((params["id"]))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var dat domain.Person
	if err := json.Unmarshal(body, &dat); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = h.service.Update(dat, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Person doesn't exist"))
		return
	}
}
