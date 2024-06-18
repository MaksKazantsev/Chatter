package utils

import (
	"fmt"
	"net/smtp"
)

type Smtp interface {
	SendCode(code, target string) error
}

const (
	sub  = "Email verification"
	body = "Code: "
)

type emailsender struct {
	smtpHost       string
	smtpPort       int
	senderEmail    string
	senderPassword string
}

func (s *emailsender) SendCode(code, target string) error {
	msg := "From: " + s.senderEmail + "\r\n" +
		"To: " + target + "\r\n" +
		"Subject: " + sub + "\r\n\r\n" +
		body + code

	auth := smtp.PlainAuth("", s.senderEmail, s.senderPassword, s.smtpHost)
	err := smtp.SendMail(fmt.Sprintf("%s:%d", s.smtpHost, s.smtpPort), auth, s.senderEmail, []string{target}, []byte(msg))
	if err != nil {
		return fmt.Errorf("failed to send an email: %v", err)
	}
	return nil
}

func NewSmtp() Smtp {
	return &emailsender{
		smtpHost:       "smtp.gmail.com",
		smtpPort:       587,
		senderEmail:    "maksimkazantsev2003@gmail.com",
		senderPassword: "enot jmup nwgt epkm",
	}
}
