package main

import (
	"EmailExchangeRate/emailHandler"
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

var (
	from   = "tempusertest9000@gmail.com"
	passwd = os.Getenv("EMAIL_PASSWD")

	smtpHost = "smtp.gmail.com"
	smtpPort = "587"
)

type MyEvent struct {
	Mess string `json:"mess"`
}

func handler(ctx context.Context, event MyEvent) {
	ms := emailHandler.MailSender{
		From:     from,
		Passwd:   passwd,
		SmtpHost: smtpHost,
		SmtpPort: smtpPort,
	}
	message := event.Mess
	if message == "" {
		message = "Hello world"
	}

	subject := "Subject: Test Email"
	emailBody := fmt.Sprintf("This is the email body: %v", message)
	to := []string{
		from,
		"pharasiuk01@gmail.com",
	}

	ms.SendMessage(subject, emailBody, to)
}

func main() {
	lambda.Start(handler)
}
