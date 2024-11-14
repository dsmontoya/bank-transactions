package transaction

import (
	"bytes"
	"html/template"
)

const (
	templateText = `
Account Balance

Total Balance: {{.TotalBalance}}
{{range .MonthlyReports}}
Month: {{.Month}}
Number of transactions: {{.TransactionCount}}
Average debit amount: {{.AverageDebit}}
Average credit amount: {{.AverageCredit}}
{{end}}
`
)

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

// Format generates a report in a human-readable format
func (r *Report) Format() (string, error) {
	totalBalance := r.TotalBalance

	data := struct {
		TotalBalance   float64
		MonthlyReports []MonthlyReport
	}{
		TotalBalance:   totalBalance,
		MonthlyReports: r.MonthlyReports,
	}

	var tpl bytes.Buffer
	t := template.Must(template.New("report").Parse(templateText))
	if err := t.Execute(&tpl, data); err != nil {
		return "", err
	}

	return tpl.String(), nil
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
