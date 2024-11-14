package transaction

type MonthlyReport struct {
	Month            string  `json:"month"`
	TransactionCount int     `json:"transaction_count"`
	AverageCredit    float64 `json:"average_credit"`
	AverageDebit     float64 `json:"average_debit"`
}

type Report struct {
	TotalBalance   float64
	MonthlyReports []MonthlyReport
}

func GenerateReport(transactions []Transaction) Report {
	monthlyData := make(map[string][]Transaction)
	totalBalance := 0.0

	for _, transaction := range transactions {
		month := transaction.Date[:7] // Assuming the date format is YYYY/MM/DD
		monthlyData[month] = append(monthlyData[month], transaction)
		totalBalance += transaction.Amount
	}

	var reports []MonthlyReport
	for month, trans := range monthlyData {
		var creditSum, debitSum float64
		var creditCount, debitCount int

		for _, t := range trans {
			if t.Amount > 0 {
				creditSum += t.Amount
				creditCount++
			} else {
				debitSum += t.Amount
				debitCount++
			}
		}

		reports = append(reports, MonthlyReport{
			Month:            month,
			TransactionCount: len(trans),
			AverageCredit:    creditSum / float64(creditCount),
			AverageDebit:     debitSum / float64(debitCount),
		})
	}

	return Report{
		TotalBalance:   totalBalance,
		MonthlyReports: reports,
	}
}
