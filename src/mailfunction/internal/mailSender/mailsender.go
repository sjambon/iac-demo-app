package mailSender

type IMailSender interface {
	SendMail(mailAddress string, subject string) error
}
