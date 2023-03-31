package internal

type Operation interface {
	Validate() error
}

type OperationInput func() (Operation, error)
