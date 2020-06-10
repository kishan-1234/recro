package controllers

import (
	"encoding/json"
	"fmt"
	"recro/models"

	"github.com/aws/aws-lambda-go/events"
)

const (
	constRequestToken  = "request_token"
	constAction        = "action"
	constStatus        = "status"
	constLogin         = "login"
	constSuccess       = "success"
	constLocation      = "Asia/Kolkata"
	constKiteApiSecret = ".kite.secret"
	constRunmode       = "runmode"
)

type UserDetails struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(request events.APIGatewayProxyRequest) string {

	returnJSON := returnJSON{}
	returnJSON.Code = "200"
	returnJSON.Msg = "Success"
	var model = make(map[string]interface{})
	userDetais := UserDetails{}

	for {

		err := json.Unmarshal([]byte(request.Body), &userDetais)
		if err != nil {
			fmt.Println("Error in unmarshalling the data", err.Error())
			returnJSON.Code = "510"
			returnJSON.Msg = "Unmarshal Error!"
			break
		}

		if userDetais.Password == "" || userDetais.Email == "" {
			returnJSON.Code = "509"
			returnJSON.Msg = "Invalid request!"
			break
		}

		user, err := models.GetUserByEmailandPassword(userDetais.Password, userDetais.Email)
		if err != nil {
			returnJSON.Code = "508"
			returnJSON.Msg = "Incorrect email or password"
			break
		}

		model[user.Name] = "HI"
		returnJSON.Model = model
		break
	}

	stringJSON, err := json.Marshal(returnJSON)
	if err != nil {
		fmt.Println("Error in marshaling response JSON struct. Struct is ", returnJSON, err.Error())
	}
	return string(stringJSON)
}
