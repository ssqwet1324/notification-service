package email

import (
	"fmt"
	"log"
	"net/smtp"
	"service_notification/internal/config"
	"strconv"
	"strings"
)

type Email struct {
	Mail     string
	Password string
	Host     string
	Port     string
	Auth     smtp.Auth
}

func NewEmail(cfg *config.Config) *Email {
	auth := smtp.PlainAuth("", cfg.SMTPEmail, cfg.SMTPPassword, cfg.SMTPHost)
	return &Email{
		cfg.SMTPEmail,
		cfg.SMTPPassword,
		cfg.SMTPHost,
		strconv.Itoa(cfg.SMTPPort),
		auth}
}

func (e *Email) Send(to, subject, message string) error {
	toMail := []string{to}
	msg := []byte("To: " + strings.Join(toMail, ",") + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + message + "\r\n")
	addr := fmt.Sprintf("%s:%s", e.Host, e.Port)
	err := smtp.SendMail(addr, e.Auth, e.Mail, toMail, msg)
	if err != nil {
		return fmt.Errorf("SendEmail Error: %v", err)
	}
	log.Printf("Send Email Success")
	return nil
}
