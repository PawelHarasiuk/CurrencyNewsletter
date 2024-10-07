package main

import (
	"EmailExchangeRate/emailHandler"
	"os"
)

var (
	from   = "tempusertest9000@gmail.com"
	passwd = os.Getenv("EMAIL_PASSWD")

	smtpHost = "smtp.gmail.com"
	smtpPort = "587"
)

func main() {
	ms := emailHandler.MailSender{
		From:     from,
		Passwd:   passwd,
		SmtpHost: smtpHost,
		SmtpPort: smtpPort,
	}

	subject := "Subject: Test Email"
	emailBody := "This is the email body."
	to := []string{
		from,
		"pharasiuk01@gmail.com",
	}

	ms.SendMessage(subject, emailBody, to)
}
