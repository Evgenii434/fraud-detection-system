package api

import (
    "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/transactions", GetTransactions).Methods("GET")
    router.HandleFunc("/transactions", CreateTransaction).Methods("POST")
    return router
}