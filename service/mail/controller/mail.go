package controller

import (
	"mail-service/pkg/log"
	"mail-service/pkg/server"
	"mail-service/service/mail/dao"
	"mail-service/service/mail/model"
	"net"
	"net/mail"
	"strings"
)

// type Email struct {
// 	To      []string `json:"to"`
// 	Subject string   `json:"subject"`
// 	Body    string   `json:"body"`
// }

func SendMail() {
	Emails := &dao.Emails{}
	err := Emails.GetEmails()
	if err != nil {
		log.Println(log.LogLevelError, "init: repo.GetEmails()", err.Error())
		return
	}

	validedEmail := []string{}
	var valided bool
	for _, email := range Emails.Emails {
		valided = ValidateEmail(email.Email)
		if valided {
			validedEmail = append(validedEmail, email.Email)
		}
	}
	email := model.CreateEmail(validedEmail, server.Config.GetString("SUBJECT"), server.Config.GetString("BODY"))

	err = email.SendEmail()
	if err != nil {
		log.Println(log.LogLevelDebug, "SendMail: email.SendEmail()", err.Error())
		return
	}
}

func ValidateEmail(email string) bool {
	// First it checks for email address format.
	_, err := mail.ParseAddress(email)
	if err != nil {
		log.Println(log.LogLevelError, "ValidateEmail: mail.ParseAddress(email)", err.Error())
		return false
	}

	//make sure that domain name is valid.
	emailPart := strings.SplitN(email, "@", 2)
	emailHost := emailPart[1]
	mx, err := net.LookupMX(emailHost)
	if err != nil {
		log.Println(log.LogLevelError, "ValidateEmail: LookupMX(emailHost)", err.Error())
		return false
	}
	if len(mx) == 0 {
		return false
	}
	return true
}
