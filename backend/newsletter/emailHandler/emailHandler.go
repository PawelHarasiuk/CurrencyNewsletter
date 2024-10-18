package emailHandler

import (
	"log/slog"
	"net/smtp"
)

type MailSender struct {
	From     string
	Passwd   string
	SmtpHost string
	SmtpPort string
}

func (ms *MailSender) SendMessage(subject, emailBody string, to []string) {
	message := []byte(subject + "\n" + emailBody)
	auth := smtp.PlainAuth("", ms.From, ms.Passwd, ms.SmtpHost)
	err := smtp.SendMail(ms.SmtpHost+":"+ms.SmtpPort, auth, ms.From, to, message)
	if err != nil {
		slog.Error(err.Error())
		return
	}
}
