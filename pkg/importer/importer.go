package importer

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"transactions/pkg/transaction"

	"go.uber.org/zap"
)

type CSVImporter struct {
	Logger *zap.Logger
	Reader io.Reader
}

func (c *CSVImporter) Import() ([]transaction.Transaction, error) {
	transactions := []transaction.Transaction{}

	r := csv.NewReader(c.Reader)

	// Ignore the first row (header)
	if _, err := r.Read(); err != nil {
		return nil, fmt.Errorf("error reading header: %w", err)
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading record: %w", err)
		}

		amount, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing amount %s: %w", record[2], err)
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("error parsing transaction id %s: %w", record[0], err)
		}

		transactions = append(transactions, transaction.Transaction{
			ID:     int64(id),
			Date:   record[1],
			Amount: amount,
		})
	}

	return transactions, nil
}
