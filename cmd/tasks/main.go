package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nrawe/tasks-kata/tasks"
)

func handle(request events.APIGatewayV2HTTPRequest) []byte {
	app := tasks.NewApp("", nil)

	result, err := app.GetTaskList(0, "", "")

	if err != nil {
		log.Printf("ERROR: application: %s\n", err)
	}

	payload, err := json.Marshal(result)

	if err != nil {
		log.Printf("ERROR: service: %s", err)
	}

	return payload
}

func main() {
	lambda.Start(handle)
}
