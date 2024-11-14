package transaction

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type Transaction struct {
	ID     string  `json:"id"`
	Date   string  `json:"date"`
	Amount float64 `json:"amount"`
}

type Handler struct {
	Logger            *zap.Logger
	TransactionWriter interface {
		Write(transactions []Transaction) error
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var transactions []Transaction
	if err := json.NewDecoder(r.Body).Decode(&transactions); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := h.TransactionWriter.Write(transactions); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	_ = GenerateReport(transactions)

	h.Logger.Info("transactions written", zap.Int("transactions_count", len(transactions)))
	w.WriteHeader(http.StatusCreated)
}
