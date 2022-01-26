package usecases

import (
	"github.com/SamuelVasconc/go-sqs-worker/interfaces"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type sqsUseCase struct {
	sqsRepository interfaces.SqsRepository
}

//NewSqsUseCase ...
func NewSqsUseCase(sqsRepository interfaces.SqsRepository) interfaces.SqsUseCase {
	return &sqsUseCase{sqsRepository}
}

func (s *sqsUseCase) GetMessages(URL string) ([]*sqs.Message, error) {
	messages, err := s.sqsRepository.GetMessages(URL)
	return messages, err
}

func (s *sqsUseCase) DeleteMessage(URL string, message *sqs.Message) error {
	err := s.sqsRepository.DeleteMessage(URL, message)
	return err
}

func (s *sqsUseCase) PublishMessage(URL string, message string) error {
	err := s.sqsRepository.PublishMessage(URL, message)
	return err
}
