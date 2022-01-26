package interfaces

import (
	"github.com/SamuelVasconc/go-sqs-worker/models"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SqsRepository interface {
	GetMessages(URL string) ([]*sqs.Message, error)
	DeleteMessage(URL string, message *sqs.Message) error
	PublishMessage(URL string, message string) error
}

type MySqlRepository interface {
	SetProtocol(status, date, protocol string, limit int) error
	GenerateProtocol() (string, error)
	GetLines(protocol string) ([]*models.Transaction, error)
	UpdateLine(obs, status string, id int64) error
}
