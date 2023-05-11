package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nataliadiasa/register/person"
)

func main() {
	repo := person.NewRepository()
	serv := person.NewService(repo)
	hand := person.NewHandler(serv)

	r := mux.NewRouter()

	r.HandleFunc("/person/{id}", hand.Get).Methods(http.MethodGet)
	r.HandleFunc("/person/{id}", hand.Update).Methods(http.MethodPut)
	r.HandleFunc("/person", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			hand.Create(w, r)
			return
		} else if r.Method == http.MethodGet {
			hand.GetAll(w, r)
			return
		}
	})

	http.Handle("/", r)
	http.ListenAndServe(":8090", nil)
}
