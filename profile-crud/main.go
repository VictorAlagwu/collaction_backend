package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func getProfile(req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var msg Response

	profileData, err := GetProfile(req)
	if err != nil {
		msg = Response{
			Message: "Error Retrieving Profile",
			Data:    "",
			Status:  http.StatusInternalServerError,
		}
		jsonData, _ := json.Marshal(msg)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       string(jsonData),
		}, err

	}

	if profileData == nil {
		msg = Response{
			Message: "no user Profile found",
			Data:    "",
			Status:  http.StatusNotFound,
		}
		jsonData, _ := json.Marshal(msg)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       string(jsonData),
		}, nil
	}

	msg = Response{
		Message: "Successfully Retrieving Profile",
		Data:    profileData,
		Status:  http.StatusOK,
	}
	jsonData, _ := json.Marshal(msg)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(jsonData),
	}, err
}

func createProfile(req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	err := CreateProfile(req)
	if err != nil {
		tmsg := Response{
			Message: fmt.Sprintf("%v", err),
			Data:    "",
			Status:  http.StatusBadRequest,
		}
		jsonData, _ := json.Marshal(tmsg)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       string(jsonData),
		}, nil

	}

	tmsg := Response{
		Message: "Profile Created",
		Data:    "",
		Status:  http.StatusCreated,
	}
	jsonData, _ := json.Marshal(tmsg)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       string(jsonData),
	}, nil
}

func updateProfile(req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	err := UpdateProfile(req)

	if err != nil {
		tmsg := Response{
			Message: fmt.Sprintf("%v", err),
			Data:    "",
			Status:  http.StatusInternalServerError,
		}
		jsonData, _ := json.Marshal(tmsg)

		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       string(jsonData),
		}, err

	}

	tmsg := Response{
		Message: "profile update successful",
		Data:    "",
		Status:  http.StatusOK,
	}
	jsonData, _ := json.Marshal(tmsg)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(jsonData),
	}, nil
}

func handler(req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	method := strings.ToLower(req.RequestContext.HTTP.Method)

	var res events.APIGatewayProxyResponse
	var err error

	switch method {
	case "get":
		res, err = getProfile(req)
	case "post":
		res, err = createProfile(req)
	case "put":
		res, err = updateProfile(req)
	default:
		jsonData, _ := json.Marshal(map[string]interface{}{"message": "Not implemented"})
		res = events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotImplemented,
			Body:       string(jsonData),
		}
	}

	return res, err
}

func main() {
	lambda.Start(handler)
}
