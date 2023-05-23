package main

import (
	"net/http"

	"github.com/gorilla/mux"
	bankaccount "github.com/nataliadiasa/register/bank_account"
	"github.com/nataliadiasa/register/person"
	"github.com/nataliadiasa/register/transaction"
)

func main() {
	personRepository := person.NewRepository()
	personService := person.NewService(personRepository)
	personHandler := person.NewHandler(personService)

	bankAccountRepository := bankaccount.NewRepository(personRepository)
	bankAccountService := bankaccount.NewService(bankAccountRepository, personService)
	bankAccountHandler := bankaccount.NewHandler(bankAccountService)

	transactionRepository := transaction.NewRepository(bankAccountRepository)
	transactionService := transaction.NewService(transactionRepository, bankAccountService)
	transactionHandler := transaction.NewHandler(transactionService)

	r := mux.NewRouter()

	r.HandleFunc("/person/{id}", personHandler.Get).Methods(http.MethodGet)
	r.HandleFunc("/person/{id}", personHandler.Update).Methods(http.MethodPut)
	r.HandleFunc("/person", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			personHandler.Create(w, r)
			return
		} else if r.Method == http.MethodGet {
			personHandler.GetAll(w, r)
			return
		}
	})

	r.HandleFunc("/account", bankAccountHandler.Create).Methods(http.MethodPost)
	r.HandleFunc("/account/{id}", bankAccountHandler.Get).Methods(http.MethodGet)

	r.HandleFunc("/transaction", transactionHandler.Create).Methods(http.MethodPost)

	http.Handle("/", r)
	http.ListenAndServe(":8090", nil)
}
