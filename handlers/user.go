package handlers

import (
	"fmt"
	"recro/controllers"
	"strings"
)

func UserHandler() (returnString string) {

	request := controllers.GetRequest()
	fmt.Println("Handling user")

	switch true {
	case strings.Contains(request.Path, "/user/all"):
		if request.HTTPMethod == "GET" {
			returnString = controllers.RetrieveAllUsers(request.QueryStringParameters)
			break
		}
	case strings.Contains(request.Path, "/user/id"):
		if request.HTTPMethod == "GET" {
			returnString = controllers.RetrieveUserById(request.QueryStringParameters)
			break
		}
	case strings.Contains(request.Path, "/user/set_password"):
		if request.HTTPMethod == "PUT" {
			returnString = controllers.SetUserPassword(request)
			break
		}
	case strings.Contains(request.Path, "/user/search"):
		if request.HTTPMethod == "POST" {
			returnString = controllers.GetUserByPhoneNumber(request)
			break
		}
	}
	return
}
