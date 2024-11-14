package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"transactions/pkg/importer"
	"transactions/pkg/transaction"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	file, err := os.Open("/assets/transactions.csv") //TODO: move to a flag
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

	resp, err := http.Post("http://transactions:8080/transactions", "application/json", bytes.NewBuffer(jsonData))
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
