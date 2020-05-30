package email

import (
	"fmt"
	"net/smtp"
)

type smtpDispacher struct {
	from     string
	password string
	host     string
	port     string
}

// Address URI to smtp server
func (s *smtpDispacher) Address() string {
	return s.host + ":" + s.port
}

func BuildSMTPDispatcher(from, password, host, port string) *smtpDispacher {
	dispatcher := smtpDispacher{
		from:     from,
		password: password,
		host:     host,
		port:     port,
	}

	return &dispatcher
}

func (dp *smtpDispacher) SendEmail(msg *Message) (ok bool, err error) {

	// Message.
	message := []byte(msg.Body) // Authentication.

	// authenticate
	auth := smtp.PlainAuth("", dp.from, dp.password, dp.host)
	fmt.Println(auth)
	// Sending email.
	err = smtp.SendMail(dp.Address(), auth, dp.from, msg.Recipient, message)
	if err != nil {
		return false, err
	}

	fmt.Println("Email Sent!")
	return true, nil
}
