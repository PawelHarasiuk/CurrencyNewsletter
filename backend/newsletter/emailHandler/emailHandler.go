package emailHandler

import (
	"EmailExchangeRate/utils"
	"fmt"
	"log/slog"
	"net/smtp"
)

type MailSender struct {
	From     string
	Passwd   string
	SmtpHost string
	SmtpPort string
}

func (ms MailSender) SendEmails(to []string) {
	subject := "Subject: Daily Currency Newsletter"
	emailBody := createMessage()
	message := []byte(subject + "\n" + emailBody)
	auth := smtp.PlainAuth("", ms.From, ms.Passwd, ms.SmtpHost)
	err := smtp.SendMail(ms.SmtpHost+":"+ms.SmtpPort, auth, ms.From, to, message)
	if err != nil {
		slog.Error(err.Error())
		return
	}
}

func createMessage() string {
	currencyTwo := "USD"
	currencyOne := "EUR"
	takeRate := utils.GetRate(currencyTwo, currencyOne)

	message := fmt.Sprintf(
		"Hello,\n\n"+
			"We are excited to share the latest updates on our transaction platform!\n\n"+
			"Rate between %s and %s is %.3f.\n\n"+
			"Thank you for choosing our services. Stay tuned for more updates.\n\n",
		currencyOne, currencyTwo, takeRate)
	return message
}
