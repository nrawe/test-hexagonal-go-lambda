package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nrawe/test-hexagonal-go-lambda/internal"
	"github.com/nrawe/test-hexagonal-go-lambda/tasks"
)

type OperationFactory struct {
	request events.APIGatewayV2HTTPRequest
}

func (f OperationFactory) getTaskList() tasks.TaskListQuery {
	return tasks.TaskListQuery{}
}

func (f OperationFactory) Get() internal.Operation {
	path := f.request.RawPath
	method := f.request.RequestContext.HTTP.Method

	switch {
	case path == "/tasks" && method == "GET":
		return f.getTaskList()
	default:
		return nil
	}
}

func handle(request events.APIGatewayV2HTTPRequest) []byte {
	app := tasks.Application{}
	operationFactory := OperationFactory{
		request: request,
	}

	result, err := app.Run(operationFactory.Get())

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
