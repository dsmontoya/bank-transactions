package main

import (
	"log"
	"net/http"
	"transactions/pkg/transaction"
)

func main() {
	handler := transaction.Handler{}
	http.HandleFunc("/transactions", handler.Handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
