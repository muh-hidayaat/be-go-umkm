package helpers

import (
	"be-go-umkm/apps/config"
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"
)

// SendEmail sends an email using SMTP
func SendEmail(to []string, subject, body string) error {
	smtpConfig := config.GetSMTPConfig()
	auth := smtp.PlainAuth("", smtpConfig.Username, smtpConfig.Password, smtpConfig.Host)
	message := formatEmailMessage(smtpConfig.From, to, subject, body)
	addr := fmt.Sprintf("%s:%s", smtpConfig.Host, smtpConfig.Port)

	if smtpConfig.Port == "465" {
		return sendEmailWithTLS(addr, smtpConfig.Host, auth, smtpConfig.From, to, message)
	}

	// Default: Send using STARTTLS
	return smtp.SendMail(addr, auth, smtpConfig.From, to, []byte(message))
}

// formatEmailMessage formats the email content
func formatEmailMessage(from string, to []string, subject, body string) string {
	return fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/plain; charset=\"utf-8\"\r\n\r\n%s",
		from, strings.Join(to, ","), subject, body)
}

// sendEmailWithTLS sends an email using a TLS connection
func sendEmailWithTLS(addr, host string, auth smtp.Auth, from string, to []string, message string) error {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		return err
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}
	defer client.Close()

	if err = authenticateAndSend(client, auth, from, to, message); err != nil {
		return err
	}

	return client.Quit()
}

// authenticateAndSend handles authentication and recipient processing
func authenticateAndSend(client *smtp.Client, auth smtp.Auth, from string, to []string, message string) error {
	if err := client.Auth(auth); err != nil {
		return err
	}
	if err := client.Mail(from); err != nil {
		return err
	}

	for _, recipient := range to {
		if err := client.Rcpt(recipient); err != nil {
			return err
		}
	}

	w, err := client.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	return w.Close()
}
