package tools

import (
	"crypto/tls"
	"os"

	gomail "gopkg.in/mail.v2"
)

// TODO: WIP, need to be fixed and tested
func SendEmail(emailTo string, subject string, mailBody string) error {
	m := gomail.NewMessage()

	smtpHost := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")

	// Set E-Mail sender
	m.SetHeader("From", from)

	// Set E-Mail receivers
	m.SetHeader("To", emailTo)

	// Set E-Mail subject
	m.SetHeader("Subject", subject)

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/html", mailBody)

	// Settings for SMTP server
	d := gomail.NewDialer(smtpHost, port, from, "")

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {

		return err
	}

	return nil
}
