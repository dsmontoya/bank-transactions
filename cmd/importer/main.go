package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("/assets/transactions.csv") //TODO: move to a flag
	if err != nil {
		log.Fatal(fmt.Errorf("error opening file: %w", err))
	}
	defer file.Close()

	r := csv.NewReader(file)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(fmt.Errorf("error reading record: %w", err))
		}

		fmt.Println(record)
	}
}
