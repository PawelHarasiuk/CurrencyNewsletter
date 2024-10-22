package main

import (
	"EmailExchangeRate/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

type Message struct {
	Email string `json:"email"`
}

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if event.Body == "" {
		fmt.Println("Request body is empty")
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Request body cannot be empty",
		}, nil
	}

	var message Message
	err := json.Unmarshal([]byte(event.Body), &message)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Error unmarshalling json",
		}, err
	}

	if event.HTTPMethod == http.MethodPost {
		err = utils.UploadEmail(message.Email)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusNotFound,
				Body:       "Incorrect email",
			}, err
		}
	} else if event.HTTPMethod == http.MethodDelete {
		err = utils.DeleteMessage(message.Email)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusNotFound,
				Body:       "Incorrect email",
			}, err
		}
	} else {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "Wrong endpoint",
		}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       "Success",
	}, err
}

func main() {
	lambda.Start(handler)
}
