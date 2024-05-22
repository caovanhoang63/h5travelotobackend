package gosmtp

import (
	"errors"
	"h5travelotobackend/email"
	"log"
	"time"
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
	//auth := smtp.PlainAuth("", g.sender, g.password, g.smtpHost)
	//err := smtp.SendMail(g.smtpHost+":"+g.smtpPort, auth, g.sender, mail.Recipient, mail.Body)
	//if err != nil {
	//	fmt.Println(err)
	//}
	time.Sleep(time.Second * 2)

	log.Println("Email send!")

	return errors.New("")
}
