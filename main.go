package main

import (
	ctx "context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"recro/controllers"
	"recro/handlers"
	"recro/models"

	"github.com/astaxie/beego/orm"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

func init() {
	setUpViper()
	setAWSenvVariables()
}

const (
	constRunmode  = "runmode"
	constLocation = "Asia/Kolkata"
)

func main() {
	lambda.Start(Handler)
}

func Handler(lambdactx ctx.Context, req map[string]interface{}) (events.APIGatewayProxyResponse, error) {

	fmt.Println("++++++++++++++++START OF REQUEST++++++++++++++++")

	headers := map[string]string{
		"Access-Control-Allow-Origin": "*",
	}

	awsResponse := events.APIGatewayProxyResponse{
		IsBase64Encoded: false,
		Body:            `{"code":"512","msg":"` + viper.GetString("maintenance.512") + `","model":{}}`,
		Headers:         headers,
		StatusCode:      200,
	}

	MAINTENANCEMODE := os.Getenv("MAINTENANCEMODE")
	if cast.ToInt(MAINTENANCEMODE) == 1 {
		return awsResponse, nil
	}

	var request events.APIGatewayProxyRequest
	requestBytes, _ := json.Marshal(req)
	_ = json.Unmarshal(requestBytes, &request)

	var returnString string

	lowerCaseHeaders(request)

	controllers.SetRequest(request)

	if registerDatabase() {
		fmt.Println(errors.New("Cannot connect to database"))
		awsResponse.StatusCode = 500
		awsResponse.Body = `{ "message": "Server is busy, please try again later" }`
		return awsResponse, nil
	}

	switch true {
	case strings.Contains(request.Path, "/user/all"):
		returnString = handlers.UserHandler()
		break
	case strings.Contains(request.Path, "/user/id"):
		returnString = handlers.UserHandler()
		break
	case strings.Contains(request.Path, "/user/set_password"):
		returnString = handlers.UserHandler()
		break
	case strings.Contains(request.Path, "/user/search"):
		returnString = handlers.UserHandler()
		break
	case strings.Contains(request.Path, "/login"):
		returnString = handlers.LoginHandler()
		break

	default:
		fmt.Println("Path did not match")
		// returnString = handlers.DefaultHandle()
		break

	}

	// attach all the data to segment before exiting the method
	awsResponse.Body = returnString
	fmt.Println("response", awsResponse)
	fmt.Println("++++++++++++++++END OF REQUEST++++++++++++++++")
	return awsResponse, nil
}

func setUpViper() {

	viper.AddConfigPath("./conf")
	viper.SetConfigName("env")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error in reading conf file through viper.Error: ", err)
	}
	viper.SetEnvPrefix("global")
	// prefix for viper global variables
}

func setAWSenvVariables() {

	setEnvVariable("MAINTENANCEMODE", "0")
	setEnvVariable("RUNMODE", "dev")

}

func setEnvVariable(key string, defaultValue string) {

	value := os.Getenv(strings.ToUpper(key))

	if value == "" {
		value = defaultValue
	}

	viper.BindEnv(strings.ToLower(key))
	os.Setenv("GLOBAL_"+strings.ToUpper(key), value)
}

func lowerCaseHeaders(info events.APIGatewayProxyRequest) {
	for i, v := range info.Headers {
		info.Headers[strings.ToLower(i)] = v
	}
}

func getRunmode() (runmode string) {

	runmode = viper.GetString(constRunmode)
	return
}

//function to register the database to beego orm
func registerDatabase() (isError bool) {
	isError = true
	runmode := getRunmode()
	//register db in do server
	mysql := viper.Get(runmode + ".mysql").(map[string]interface{})
	fmt.Println(mysql)
	mysqlConf := mysql["user"].(string) + ":" + mysql["password"].(string) + "@tcp(" + mysql["host"].(string) + ")/" + mysql["database"].(string) + "?interpolateParams=true&charset=utf8&parseTime=True&loc=Asia%2FKolkata"

	for breaker := 5; breaker > 0; breaker-- {

		_, err := orm.GetDB()
		if err != nil {

			fmt.Println("conf", mysqlConf)
			if err := orm.RegisterDataBase("default", "mysql", mysqlConf); err != nil {
				fmt.Println("Could not register Database!.Error: ", err)
			}
		}

		orm.Debug = true

		err = models.MysqlTest()
		if err != nil {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("^^^^^^^^^^^^^mysql connection error ^^^^^^^^^^^^^^")
			continue
		} else {
			fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$Connected to MySQL Server$$$$$$$$$$$$$$$$$$$$$$$$$")
		}
		orm.DefaultTimeLoc, _ = time.LoadLocation(constLocation)
		isError = false
		break
	}
	return
}
