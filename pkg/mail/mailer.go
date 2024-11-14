package mail

import "gopkg.in/gomail.v2"

type smtpConfig struct {
	host string
	port int
	user string
	pass string
}

type Mailer struct {
	from   string
	dialer gomail.Dialer
}

type MailerOption func(*Mailer)

func WithSMTPAddr(host string, port int) MailerOption {
	return func(m *Mailer) {
		m.dialer.Host = host
		m.dialer.Port = port
	}
}

func WithFrom(from string) MailerOption {
	return func(m *Mailer) {
		m.from = from
	}
}

func NewMailer(opts ...MailerOption) *Mailer {
	m := &Mailer{}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

func (m *Mailer) Send(to, subject, body string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", m.from)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", body)

	return m.dialer.DialAndSend(message)
}
