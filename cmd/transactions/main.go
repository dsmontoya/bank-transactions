package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"transactions/pkg/mail"
	"transactions/pkg/transaction"

	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	mailer := mail.NewMailer(
		mail.WithFrom("no-reply@example.com"),
		mail.WithSMTPAddr("localhost", 587),
	)

	handler := transaction.Handler{
		Logger:            logger,
		TransactionWriter: transaction.NewWriter(conn),
		Notifier:          &transaction.Notifier{Logger: logger, Mailer: mailer},
	}

	http.HandleFunc("/transactions", handler.Handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
