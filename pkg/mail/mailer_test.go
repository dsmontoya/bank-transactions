package mail

import (
	"fmt"
	"net"
	"testing"
)

func TestMailer_Send(t *testing.T) {
	mailer := NewMailer(WithSMTPAddr("127.0.0.1", 2525), WithFrom("no-reply@example.com"))
	err := mailer.Send("to@example.com", "Subject", "Body")
	if err != nil {
		t.Fatalf("Send() failed: %v", err)
	}
}

func startSMTPServer() (net.Listener, error) {
	ln, err := net.Listen("tcp", "127.0.0.1:2525")
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				return
			}
			go handleSMTPConnection(conn)
		}
	}()

	return ln, nil
}

func handleSMTPConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Fprintln(conn, "220 Welcome")
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		cmd := string(buf[:n])
		switch {
		case cmd == "QUIT\r\n":
			fmt.Fprintln(conn, "221 Bye")
			return
		case cmd == "HELO localhost\r\n":
			fmt.Fprintln(conn, "250 Hello localhost")
		case cmd == "MAIL FROM:<no-reply@example.com>\r\n":
			fmt.Fprintln(conn, "250 OK")
		case cmd == "RCPT TO:<to@example.com>\r\n":
			fmt.Fprintln(conn, "250 OK")
		case cmd == "DATA\r\n":
			fmt.Fprintln(conn, "354 End data with <CR><LF>.<CR><LF>")
		default:
			if cmd == ".\r\n" {
				fmt.Fprintln(conn, "250 OK")
			} else {
				fmt.Fprintln(conn, "250 OK")
			}
		}
	}
}

func TestMain(m *testing.M) {
	ln, err := startSMTPServer()
	if err != nil {
		fmt.Printf("Failed to start SMTP server: %v\n", err)
		return
	}
	defer ln.Close()

	m.Run()
}
