package controllers

import (
	"time"

	"github.com/aws/aws-lambda-go/events"
)

var _request = events.APIGatewayProxyRequest{}

type returnJSON struct {
	Code  string      `json:"code"`
	Msg   string      `json:"msg"`
	Model interface{} `json:"model"`
}

func SetRequest(request events.APIGatewayProxyRequest) {
	_request = request
}

func GetRequest() events.APIGatewayProxyRequest {
	return _request
}

func getCurrentIndianTime() (currentTime time.Time) {

	loc, _ := time.LoadLocation(constLocation)
	currentTime = time.Now().In(loc)
	return
}
