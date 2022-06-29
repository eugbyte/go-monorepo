package models

import (
	"log"

	colors "github.com/TwinProduction/go-color"
	"github.com/aws/aws-lambda-go/events"
)

type Response = events.APIGatewayProxyResponse

type HttpError struct {
	StatusCode int
	Err        error
}

func (er HttpError) Error() string {
	return er.Err.Error()
}

func (er HttpError) Log() {
	// calling er.Err will result in stack overflow
	errorMessage := er.Err.Error()
	log.Println(colors.Red, errorMessage, colors.Reset)
}
