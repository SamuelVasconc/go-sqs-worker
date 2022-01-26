package interfaces

import (
	"github.com/SamuelVasconc/go-sqs-worker/models"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SqsUseCase interface {
	GetMessages(URL string) ([]*sqs.Message, error)
	DeleteMessage(URL string, message *sqs.Message) error
}

type MySqlUseCase interface {
	SetProtocol(status, date string, limit int) (string, error)
	GetLines(parameters models.Parameter) error
}
