package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
)

type MyEvent struct {
	// sample properties below
	InstanceId string `json:"InstanceId"`
	Status     string `json:"Status"`
}

func HandleRequest(context context.Context, ev MyEvent) (events.APIGatewayProxyResponse, error) {

	res := events.APIGatewayProxyResponse{
		IsBase64Encoded: false,
		StatusCode:      200,
		Headers:         nil,
		Body:            "Success!",
	}

	config.LoadDefaultConfig(context)

	return res, nil
}

// entry point to your lambda
func main() {
	lambda.Start(HandleRequest)
}
