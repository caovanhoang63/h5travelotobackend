package email

import (
	"bytes"
	"fmt"
	"h5travelotobackend/common"
	"html/template"
	"log"
)

type MailSender interface {
	Send(Mail) error
}

type Mail struct {
	Sender    string
	Recipient []string
	Subject   string
	Body      []byte
	Password  string
}

func NewEmail(sender, password, subject string, body []byte, recipient []string) *Mail {
	return &Mail{
		Sender:    sender,
		Recipient: recipient,
		Subject:   subject,
		Body:      body,
		Password:  password,
	}
}

func NewRecoverPasswordMail(recipient string, pinCode string) *Mail {
	RecoverIcon := "https://d3jwhct9rpti9n.cloudfront.net/room_images/854725495.png"
	subject := "Yêu cầu tạo lại mật khẩu"
	t, err := template.ParseFiles("./email/static/reset-password-mail.html")
	if err != nil {
		log.Println(err)
		return nil
	}

	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", subject, mimeHeaders)))
	t.Execute(&body, struct {
		PinCode     string
		RecoverIcon string
	}{
		PinCode:     pinCode,
		RecoverIcon: RecoverIcon,
	})
	to := []string{
		recipient,
	}
	return &Mail{
		Recipient: to,
		Subject:   subject,
		Body:      body.Bytes(),
	}
}

func NewConfirmEmail(recipient string, bk common.ConfirmBooking) *Mail {
	subject := "Đặt phòng thành công"
	t, err := template.ParseFiles("./email/static/confirm-email.html")
	if err != nil {
		log.Println(err)
		return nil
	}
	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", subject, mimeHeaders)))
	t.Execute(&body, bk)
	to := []string{
		recipient,
	}
	return &Mail{
		Recipient: to,
		Subject:   subject,
		Body:      body.Bytes(),
	}
}
