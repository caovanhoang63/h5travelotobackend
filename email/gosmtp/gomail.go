package gosmtp

import (
	"fmt"
	"h5travelotobackend/email"
	"log"
	"net/smtp"
)

type GoMail struct {
	smtpHost string
	smtpPort string
	sender   string
	password string
}

func NewGoMail(smtpHost, smtpPort, sender, password string) *GoMail {
	return &GoMail{
		smtpHost: smtpHost,
		smtpPort: smtpPort,
		sender:   sender,
		password: password,
	}
}

func (g *GoMail) Send(mail email.Mail) error {
	auth := smtp.PlainAuth("", g.sender, g.password, g.smtpHost)
	err := smtp.SendMail(g.smtpHost+":"+g.smtpPort, auth, g.sender, mail.Recipient, mail.Body)
	if err != nil {
		fmt.Println(err)
	}
	log.Println("Email sent")
	return nil
}
