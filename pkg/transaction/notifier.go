package transaction

import (
	"context"

	"go.uber.org/zap"
)

type Notifier struct {
	Logger *zap.Logger
}

func (n *Notifier) Notify(ctx context.Context, transactions []Transaction) error {
	_ = GenerateReport(transactions)
	return nil
}
