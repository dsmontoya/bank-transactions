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

func (n *Notifier) Notify(_ context.Context, transactions []Transaction) error {
	report := GenerateReport(transactions)

	body, err := report.Format()
	if err != nil {
		return err
	}

	if err := n.Mailer.Send("", "", body); err != nil {
		n.Logger.Error("failed to send email", zap.Error(err))
		return err
	}
	return nil
}
