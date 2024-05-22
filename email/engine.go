package email

type engine struct {
	mailSender MailSender
	queue      chan Mail
}

func (e *engine) Send(mail Mail) {

}

func (e *engine) Start() {

}
