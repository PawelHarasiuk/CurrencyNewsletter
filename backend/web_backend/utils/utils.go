package utils

import (
	"bufio"
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"strings"
)

var (
	bucket = "newsletter-bucket-go"
	key    = "emails.txt"
)

func DeleteMessage(email string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("eu-central-1"))
	if err != nil {
		return err
	}

	client := s3.NewFromConfig(cfg)

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	result, err := client.GetObject(context.TODO(), input)
	if err != nil {
		return err
	}
	defer result.Body.Close()

	var buffer bytes.Buffer
	scanner := bufio.NewScanner(result.Body)

	for scanner.Scan() {
		line := scanner.Text()
		if line != email {
			buffer.WriteString(line + "\n")
		}
	}

	reader := strings.NewReader(buffer.String())
	putInput := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   reader,
	}

	_, err = client.PutObject(context.TODO(), putInput)
	if err != nil {
		log.Fatalf("unable to upload updated file %v, %v", key, err)
	}
	return nil
}

func UploadEmail(email string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("eu-central-1"))
	if err != nil {
		return err
	}

	client := s3.NewFromConfig(cfg)

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	result, err := client.GetObject(context.TODO(), input)
	if err != nil {
		return err
	}
	defer result.Body.Close()

	var buffer bytes.Buffer
	scanner := bufio.NewScanner(result.Body)

	for scanner.Scan() {
		line := scanner.Text()
		buffer.WriteString(line + "\n")
	}

	buffer.WriteString(email + "\n")

	reader := strings.NewReader(buffer.String())
	putInput := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   reader,
	}

	_, err = client.PutObject(context.TODO(), putInput)
	if err != nil {
		log.Fatalf("unable to upload updated file %v, %v", key, err)
	}
	return nil
}
