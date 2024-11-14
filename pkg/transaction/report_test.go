package transaction_test

import (
	"reflect"
	"testing"
	"transactions/pkg/transaction"
)

func TestGenerateReport(t *testing.T) {
	tests := []struct {
		name         string
		transactions []transaction.Transaction
		want         transaction.Report
	}{
		{
			name: "single month transactions",
			transactions: []transaction.Transaction{
				{Date: "2023/01/01", Amount: 100.0},
				{Date: "2023/01/15", Amount: -50.0},
				{Date: "2023/01/20", Amount: 200.0},
			},
			want: transaction.Report{
				TotalBalance: 250.0,
				MonthlyReports: []transaction.MonthlyReport{
					{
						Month:            "2023/01",
						TransactionCount: 3,
						AverageCredit:    150.0,
						AverageDebit:     -50.0,
					},
				},
			},
		},
		{
			name: "multiple months transactions",
			transactions: []transaction.Transaction{
				{Date: "2023/01/01", Amount: 100.0},
				{Date: "2023/01/15", Amount: -50.0},
				{Date: "2023/02/01", Amount: 200.0},
				{Date: "2023/02/15", Amount: -100.0},
			},
			want: transaction.Report{
				TotalBalance: 150.0,
				MonthlyReports: []transaction.MonthlyReport{
					{
						Month:            "2023/01",
						TransactionCount: 2,
						AverageCredit:    100.0,
						AverageDebit:     -50.0,
					},
					{
						Month:            "2023/02",
						TransactionCount: 2,
						AverageCredit:    200.0,
						AverageDebit:     -100.0,
					},
				},
			},
		},
		{
			name:         "no transactions",
			transactions: []transaction.Transaction{},
			want: transaction.Report{
				TotalBalance:   0.0,
				MonthlyReports: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := transaction.GenerateReport(tt.transactions)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateReport() = %v, want %v", got, tt.want)
			}
		})
	}
}