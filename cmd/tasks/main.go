package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nrawe/test-hexagonal-go-lambda/tasks"
)

type HttpGatewayOperationAdapter struct {
	request events.APIGatewayV2HTTPRequest
}

func (a HttpGatewayOperationAdapter) GetOperationDetails() (tasks.OperationId, interface{}, error) {
	path := a.request.RawPath
	method := a.request.RequestContext.HTTP.Method

	switch {
	case path == "/tasks" && method == "GET":
		return a.getUserTaskList()
	default:
		return tasks.OpUnknown, nil, nil
	}
}

func (a HttpGatewayOperationAdapter) getUserTaskList() (tasks.OperationId, tasks.GetUserTaskList, error) {
	op := tasks.GetUserTaskList{}

	return tasks.OpGetUserTaskList, op, nil
}

func handle() (string, error) {
	return "Hello world", nil
}

func main() {
	lambda.Start(handle)
}
