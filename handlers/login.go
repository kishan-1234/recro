package handlers

import (
	"fmt"
	"recro/controllers"
	"strings"
)

func LoginHandler() (returnString string) {

	request := controllers.GetRequest()
	fmt.Println("Handling order")

	switch true {
	case strings.Contains(request.Path, "/login"):
		if request.HTTPMethod == "POST" {
			returnString = controllers.LoginUser(request)
			break
		}
	}
	return
}
