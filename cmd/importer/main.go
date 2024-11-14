package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"transactions/pkg/importer"
	"transactions/pkg/transaction"

	"go.uber.org/zap"
)

var (
	transactionsAddr = flag.String("transactions-addr", "transactions:8080", "Address of the transactions service")
	csvFile          = flag.String("csv-file", "/assets/transactions.csv", "Path to the CSV file")
)

func main() {
	flag.Parse()

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	file, err := os.Open(*csvFile)
	if err != nil {
		logger.Fatal(fmt.Sprintf("error opening file: %s", err))
	}
	defer file.Close()

	csvImporter := importer.CSVImporter{
		Logger: logger,
		Reader: file,
	}
	transactions, err := csvImporter.Import()
	if err != nil {
		logger.Fatal(fmt.Sprintf("error importing transactions: %s", err))
	}
	if err := sendTransactions(transactions); err != nil {
		logger.Fatal(fmt.Sprintf("error sending transactions: %s", err))
	}

	logger.Info("transactions imported")
}

func sendTransactions(transactions []transaction.Transaction) error {
	jsonData, err := json.Marshal(transactions)
	if err != nil {
		return fmt.Errorf("error marshalling transactions: %w", err)
	}

	resp, err := http.Post(fmt.Sprintf("http://%s/transactions", *transactionsAddr), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error posting transactions: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	return nil
}
