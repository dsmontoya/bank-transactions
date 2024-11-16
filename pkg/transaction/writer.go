package transaction

import (
	"context"
	"time"
	"transactions/pkg/generated/sql/transactionssql"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

type Writer struct {
	conn    *pgx.Conn
	queries *transactionssql.Queries
}

func NewWriter(conn *pgx.Conn) *Writer {
	return &Writer{
		queries: transactionssql.New(conn),
	}
}

func (w *Writer) Write(ctx context.Context, transactions []Transaction) error {
	sqlParams, err := transactionsToSQLParams(transactions)
	if err != nil {
		return err
	}
	_, err = w.queries.AddTransactions(ctx, sqlParams)
	return err
}

func transactionsToSQLParams(transactions []Transaction) (transactionssql.AddTransactionsParams, error) {
	sqlParams := transactionssql.AddTransactionsParams{}
	for _, transaction := range transactions {
		sqlParams.Ids = append(sqlParams.Ids, transaction.ID)

		amountNumeric := pgtype.Numeric{}
		amountNumeric.Set(transaction.Amount)
		sqlParams.Amounts = append(sqlParams.Amounts, amountNumeric)

		date, err := time.Parse("2006/01/02", transaction.Date)
		if err != nil {
			return transactionssql.AddTransactionsParams{}, err
		}
		sqlParams.Dates = append(sqlParams.Dates, date)

		sqlParams.UserIds = append(sqlParams.UserIds, transaction.UserID)
	}

	return sqlParams, nil
}
