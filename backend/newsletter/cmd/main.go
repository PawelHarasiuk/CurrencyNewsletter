package main

import (
	"EmailExchangeRate/emailHandler"
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
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
	message := "Hello"

	subject := "Subject: Test Email"
	emailBody := fmt.Sprintf("This is the email body: %v", message)
	to, err := ReadMailsFromS3()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(to)
	ms.SendMessage(subject, emailBody, to)
}

func main() {
	lambda.Start(handler)
}

func ReadMailsFromS3() ([]string, error) {
	bucket := "newsletter-bucket-go"
	key := "emails.txt"
	var to []string

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("eu-central-1"))
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	result, err := client.GetObject(context.TODO(), input)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	scanner := bufio.NewScanner(result.Body)
	for scanner.Scan() {
		to = append(to, scanner.Text())
	}

	if len(to) == 0 {
		return nil, errors.New("emails not found in storage")
	}

	return to, nil
}
