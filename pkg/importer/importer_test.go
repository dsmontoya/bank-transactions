package importer_test

import (
	"reflect"
	"strings"
	"testing"
	"transactions/pkg/importer"
	"transactions/pkg/transaction"

	"go.uber.org/zap/zaptest"
)

func TestCSVImporter_Import(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []transaction.Transaction
		wantErr bool
	}{
		{
			name: "valid input",
			input: `ID,Date,Amount
1,2024-01-01,+100.50
2,2024-01-02,-200.75
`,
			want: []transaction.Transaction{
				{ID: 1, Date: "2024-01-01", Amount: 100.50, UserID: 1},
				{ID: 2, Date: "2024-01-02", Amount: -200.75, UserID: 1},
			},
			wantErr: false,
		},
		{
			name: "invalid amount",
			input: `ID,Date,Amount
1,2023-01-01,invalid
`,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "empty input",
			input:   ``,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := zaptest.NewLogger(t)
			reader := strings.NewReader(tt.input)
			importer := &importer.CSVImporter{
				Logger: logger,
				Reader: reader,
			}

			transactions, err := importer.Import()
			if (err != nil) != tt.wantErr {
				t.Errorf("want error: %v, got: %v", tt.wantErr, err)
			}

			if !tt.wantErr && !reflect.DeepEqual(transactions, tt.want) {
				t.Errorf("want: %v, got: %v", tt.want, transactions)
			}
		})
	}
}
