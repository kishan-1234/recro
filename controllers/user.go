package controllers

import (
	"encoding/json"
	"fmt"
	"recro/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/spf13/cast"
)

type PasswordDetails struct {
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
}

func RetrieveAllUsers(QueryStringParameters map[string]string) string {

	returnJSON := returnJSON{}
	returnJSON.Code = "200"
	returnJSON.Msg = "Success"
	var model = make(map[string]interface{})

	for {

		users, err := models.GetAllUsers()
		if err != nil {
			returnJSON.Code = "510"
			returnJSON.Msg = "Error getting all users from DB!"
			break
		}

		for _, user := range users {
			model[cast.ToString(user.Id)] = user.Name
		}

		returnJSON.Model = model
		break
	}

	stringJSON, err := json.Marshal(returnJSON)
	if err != nil {
		fmt.Println("Error in marshaling response JSON struct. Struct is ", returnJSON, err.Error())
	}
	return string(stringJSON)
}

func RetrieveUserById(QueryStringParameters map[string]string) string {

	returnJSON := returnJSON{}
	returnJSON.Code = "200"
	returnJSON.Msg = "Success"
	var model = make(map[string]interface{})
	var id string
	var ok bool

	for {

		if id, ok = QueryStringParameters["id"]; !ok {
			returnJSON.Code = "509"
			returnJSON.Msg = "Id Query Param not found!"
			returnJSON.Model = model
			break
		}

		user, err := models.GetUserById(cast.ToInt(id))
		if err != nil {
			returnJSON.Code = "510"
			returnJSON.Msg = "Error getting user from DB!"
			break
		}

		model[cast.ToString(user.Id)] = user.Name

		returnJSON.Model = model
		break
	}

	stringJSON, err := json.Marshal(returnJSON)
	if err != nil {
		fmt.Println("Error in marshaling response JSON struct. Struct is ", returnJSON, err.Error())
	}
	return string(stringJSON)
}

func SetUserPassword(request events.APIGatewayProxyRequest) string {

	returnJSON := returnJSON{}
	returnJSON.Code = "200"
	returnJSON.Msg = "Success"
	var model = make(map[string]interface{})
	passwordDetais := PasswordDetails{}

	for {

		err := json.Unmarshal([]byte(request.Body), &passwordDetais)
		if err != nil {
			fmt.Println("Error in unmarshalling the data", err.Error())
			returnJSON.Code = "510"
			returnJSON.Msg = "Unmarshal Error!"
			break
		}

		if passwordDetais.Password == "" || passwordDetais.PhoneNumber == "" {
			returnJSON.Code = "509"
			returnJSON.Msg = "Invalid request!"
			break
		}

		user, err := models.GetUserByPhoneNumber(passwordDetais.PhoneNumber)
		if err != nil {
			returnJSON.Code = "508"
			returnJSON.Msg = "Error getting user from DB by phoneNumber!"
			break
		}

		user.Password = passwordDetais.Password
		err = models.UpdateRowByColumns(user, "updated", "password")
		if err != nil {
			fmt.Println("Error in updating user password ")
			returnJSON.Code = "507"
			returnJSON.Msg = "Error in updating user password!"
			break
		}

		returnJSON.Model = model
		break
	}

	stringJSON, err := json.Marshal(returnJSON)
	if err != nil {
		fmt.Println("Error in marshaling response JSON struct. Struct is ", returnJSON, err.Error())
	}
	return string(stringJSON)
}

func GetUserByPhoneNumber(request events.APIGatewayProxyRequest) string {

	returnJSON := returnJSON{}
	returnJSON.Code = "200"
	returnJSON.Msg = "Success"
	var model = make(map[string]interface{})

	passwordDetais := PasswordDetails{}

	for {

		err := json.Unmarshal([]byte(request.Body), &passwordDetais)
		if err != nil {
			fmt.Println("Error in unmarshalling the data", err.Error())
			returnJSON.Code = "510"
			returnJSON.Msg = "Unmarshal Error!"
			break
		}

		if passwordDetais.PhoneNumber == "" {
			returnJSON.Code = "509"
			returnJSON.Msg = "Invalid request!"
			break
		}

		user, err := models.GetUserByPhoneNumber(passwordDetais.PhoneNumber)
		if err != nil {
			returnJSON.Code = "508"
			returnJSON.Msg = "Error getting user from DB by phoneNumber!"
			break
		}

		model[cast.ToString(user.Id)] = user.Name

		returnJSON.Model = model
		break
	}

	stringJSON, err := json.Marshal(returnJSON)
	if err != nil {
		fmt.Println("Error in marshaling response JSON struct. Struct is ", returnJSON, err.Error())
	}
	return string(stringJSON)
}
