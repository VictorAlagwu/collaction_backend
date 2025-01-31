package utils

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func GetMessageHttpResponse(statusCode int, msg string) events.APIGatewayProxyResponse {
	// "Cannot go wrong"
	jsonPayload, _ := json.Marshal(map[string]interface{}{"message": msg})
	return events.APIGatewayProxyResponse{
		Body:       string(jsonPayload),
		StatusCode: statusCode,
	}
}
