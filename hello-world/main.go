package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	body := "Content"
	response, _ := json.Marshal(&body)

	return events.APIGatewayProxyResponse{
		Body:       string(response),
		StatusCode: 200,
		IsBase64Encoded: true,
		Headers: map[string]string{

		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
