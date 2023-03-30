package tasks

type TaskId string

type OperationId byte

const (
	OpUnknown OperationId = iota
	OpGetUserTaskList
)

type GetUserTaskList struct {
	after  TaskId
	before TaskId
	limit  int8
}

type OperationAdapter interface {
	GetOperationDetails() (OperationId, interface{}, error)
}

type TaskHexagon struct {
	operation OperationAdapter
}

func (t TaskHexagon) Run() {
	id, op, err := t.operation.GetOperationDetails()

	if err != nil {
		panic(err)
	}

	switch id {
	case OpGetUserTaskList:
		t.getUserTaskList(op.(GetUserTaskList))
	}
}

func (t TaskHexagon) getUserTaskList(op GetUserTaskList) {

}
