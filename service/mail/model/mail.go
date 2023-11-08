package model

import (
	"net/smtp"
)

const (
	_host                = "smtp.gmail.com"
	_address             = "smtp.gmail.com:587"
	_senderEmail         = "nguyencaohiep.work@gmail.com" // FROM
	_applicationPassword = "enmehzlwhyvrqvkv"             //PASSWORD
)

var _defaultAuth = smtp.PlainAuth("", _senderEmail, _applicationPassword, _host)

type Email struct {
	to      []string
	subject string
	body    string
}

func CreateEmail(to []string, subject string, body string) *Email {
	return &Email{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (email *Email) SendEmail() error {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := "Subject:" + email.subject + "\n" + mime + "\n" + email.body
	err := smtp.SendMail(_address, _defaultAuth, _senderEmail, email.to, []byte(msg))
	return err
}
