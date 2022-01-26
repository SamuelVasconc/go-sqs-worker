package repositories

import (
	"github.com/SamuelVasconc/go-sqs-worker/interfaces"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type sqsRepository struct {
	SQS *sqs.SQS
}

//NewSqsRepository ...
func NewSqsRepository(SQS *sqs.SQS) interfaces.SqsRepository {
	return &sqsRepository{SQS}
}

func (s *sqsRepository) GetMessages(URL string) ([]*sqs.Message, error) {
	result, err := s.SQS.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            &URL,
		MaxNumberOfMessages: aws.Int64(10),
		VisibilityTimeout:   aws.Int64(60), // 60 seconds
		WaitTimeSeconds:     aws.Int64(0),
	})
	if err != nil {
		return nil, err
	}

	return result.Messages, nil
}

func (s *sqsRepository) DeleteMessage(URL string, message *sqs.Message) error {
	_, err := s.SQS.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      &URL,
		ReceiptHandle: message.ReceiptHandle,
	})

	return err
}

func (s *sqsRepository) PublishMessage(URL string, message string) error {
	_, err := s.SQS.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    &URL,
		MessageBody: aws.String(message),
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"Squad": {
				DataType:    aws.String("String"),
				StringValue: aws.String("Middleware"),
			},
			"ApplicationName": {
				DataType:    aws.String("String"),
				StringValue: aws.String("go-sqs-worker"),
			},
		},
	})

	return err
}
