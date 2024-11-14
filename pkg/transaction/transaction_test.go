package transaction

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockTransactionWriter func(transaction Transaction) error

func (m MockTransactionWriter) Write(transaction Transaction) error {
	if m == nil {
		panic("unimplemented")
	}
	return m(transaction)
}

func TestHandle(t *testing.T) {
	tests := []struct {
		name                  string
		method                string
		body                  interface{}
		transactionWriterFunc func(transaction Transaction) error
		expectedStatusCode    int
	}{
		{
			name:                  "Invalid method",
			method:                http.MethodGet,
			body:                  nil,
			transactionWriterFunc: nil,
			expectedStatusCode:    http.StatusMethodNotAllowed,
		},
		{
			name:                  "Invalid body",
			method:                http.MethodPost,
			body:                  "invalid",
			transactionWriterFunc: nil,
			expectedStatusCode:    http.StatusBadRequest,
		},
		{
			name:   "Write error",
			method: http.MethodPost,
			body: Transaction{
				ID:     "1",
				Date:   "2023-10-01",
				Amount: 100.0,
			},
			transactionWriterFunc: func(transaction Transaction) error {
				return errors.New("write error")
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:   "Successful write",
			method: http.MethodPost,
			body: Transaction{
				ID:     "1",
				Date:   "2023-10-01",
				Amount: 100.0,
			},
			transactionWriterFunc: func(transaction Transaction) error {
				return nil
			},
			expectedStatusCode: http.StatusCreated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var bodyBytes []byte
			if tt.body != nil {
				bodyBytes, _ = json.Marshal(tt.body)
			}

			req := httptest.NewRequest(tt.method, "/transactions", bytes.NewBuffer(bodyBytes))
			rec := httptest.NewRecorder()

			handler := &Handler{
				TransactionWriter: MockTransactionWriter(tt.transactionWriterFunc),
			}

			handler.Handle(rec, req)

			if rec.Code != tt.expectedStatusCode {
				t.Errorf("expected status code %d, got %d", tt.expectedStatusCode, rec.Code)
			}
		})
	}
}