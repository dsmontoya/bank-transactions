package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

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

	r := csv.NewReader(file)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Fatal(fmt.Sprintf("error reading record: %s", err))
		}

		logger.Info("transactions imported", zap.Any("record", record))
	}
}
