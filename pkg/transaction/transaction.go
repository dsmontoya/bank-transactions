package transaction

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

const dummyEmail = "test@test.com"

type Transaction struct {
	ID     int64   `json:"id"`
	Date   string  `json:"date"`
	Amount float64 `json:"amount"`
	UserID int64   `json:"user_id"`
}

type Handler struct {
	Logger            *zap.Logger
	TransactionWriter interface {
		Write(ctx context.Context, transactions []Transaction) error
	}
	Notifier interface {
		Notify(ctx context.Context, to string, transactions []Transaction) error
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

	if err := h.TransactionWriter.Write(r.Context(), transactions); err != nil {
		h.Logger.Error("Failed to write transaction", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	go func(ctx context.Context) {
		if err := h.Notifier.Notify(ctx, dummyEmail, transactions); err != nil {
			h.Logger.Error("error notifying", zap.Error(err))
		}
	}(r.Context())

	h.Logger.Info("transactions written", zap.Int("transactions_count", len(transactions)))
	w.WriteHeader(http.StatusCreated)
}
