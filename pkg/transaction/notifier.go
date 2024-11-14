package transaction

import (
	"context"
	"transactions/pkg/mail"

	"go.uber.org/zap"
)

type Notifier struct {
	Logger *zap.Logger
	Mailer mail.Mailer
}

func (n *Notifier) Notify(ctx context.Context, transactions []Transaction) error {
	_ = GenerateReport(transactions)
	return nil
}
