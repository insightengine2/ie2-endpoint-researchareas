package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/insightengine2/ie2-endpoint-researchareas/lib"
)

func fillErrorResponse(res *events.APIGatewayProxyResponse, e error) {

	if res != nil && e != nil {
		res.Body = e.Error()
		res.StatusCode = 500
	}
}

func HandleRequest(context context.Context, ev any) (events.APIGatewayProxyResponse, error) {

	res := events.APIGatewayProxyResponse{
		IsBase64Encoded: false,
		StatusCode:      200,
		Headers:         nil,
		Body:            "Success!",
	}

	_, err := config.LoadDefaultConfig(context)

	if err != nil {
		fillErrorResponse(&res, err)
		return res, err
	}

	researchareas, err := lib.GetResearchAreas()

	if err != nil {
		fillErrorResponse(&res, err)
		return res, err
	}

	jsonRes, err := json.Marshal(researchareas)

	if err != nil {
		fillErrorResponse(&res, err)
		return res, err
	}

	res.Body = string(jsonRes)

	return res, nil
}

// entry point to your lambda
func main() {
	lambda.Start(HandleRequest)
}
